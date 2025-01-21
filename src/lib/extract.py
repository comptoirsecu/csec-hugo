import yaml
import re
import json
import os
from datetime import date, datetime

class CustomJSONEncoder(json.JSONEncoder):
    def default(self, obj):
        if isinstance(obj, (date, datetime)):
            return obj.isoformat()
        return super().default(obj)

#directory = 'csec-hugo/content/podcast/'
directory = 'csec-hugo/content/sechebdo/'
output_file = 'all_sechebdo.json'
songs_file = 'csec-hugo/data/songs.yml'
authors_file = 'csec-hugo/data/authors.yml'
guests_file = 'csec-hugo/data/guests.yml'

with open(songs_file) as songs_stream:
    try:
        songs_data = yaml.load(songs_stream, Loader=yaml.Loader)

        with open(authors_file) as authors_stream:
            authors_data = yaml.load(authors_stream, Loader=yaml.Loader)

            with open(guests_file) as guests_stream:
                guests_data = yaml.load(guests_stream, Loader=yaml.Loader)

                results = []

                for filename in os.listdir(directory):
                    if filename.endswith('.md'):
                        file_path = os.path.join(directory, filename)
                        with open(file_path) as stream:
                            file_content = stream.read()
                            parts = file_content.split('---')
                            yaml_content = parts[1]
                            markdown_content = '---'.join(parts[2:])
                            parsed_yaml = yaml.load(yaml_content, Loader=yaml.Loader)

                            if 'songs' in parsed_yaml and isinstance(parsed_yaml['songs'], list):
                                for i, song_key in enumerate(parsed_yaml['songs']):
                                    if song_key in songs_data:
                                        parsed_yaml['songs'][i] = songs_data[song_key]

                            if 'authors' in parsed_yaml and isinstance(parsed_yaml['authors'], list):
                                for i, author_key in enumerate(parsed_yaml['authors']):
                                    if author_key in authors_data:
                                        parsed_yaml['authors'][i] = authors_data[author_key]

                            if 'guests' in parsed_yaml and isinstance(parsed_yaml['guests'], list):
                                for i, guest_key in enumerate(parsed_yaml['guests']):
                                    if guest_key in guests_data:
                                        parsed_yaml['guests'][i] = guests_data[guest_key]

                            # Extract description before the {{< string
                            description_match = re.search(r'(.*?)\{\{<', markdown_content, re.DOTALL)
                            description = description_match.group(1).strip() if description_match else markdown_content.strip()

                            # Extract HTTP links with their titles and URLs
                            links = re.findall(r'\[([^\]]+)\]\((http[^\)]+)\)', markdown_content)

                            result = {
                                "metadata": parsed_yaml,
                                "content": {
                                    "description": description,
                                    "links": [{"title": title, "url": url} for title, url in links]
                                }
                            }

                            results.append(result)

                # Normalize date values to datetime for sorting
                def normalize_date(d):
                    if isinstance(d, date) and not isinstance(d, datetime):
                        return datetime.combine(d, datetime.min.time())
                    if d.tzinfo is not None:
                        return d.astimezone(tz=None).replace(tzinfo=None)
                    return d

                # Sort results by metadata date
                sorted_results = sorted(results, key=lambda x: normalize_date(x['metadata'].get('date', datetime.min)), reverse=True)

                result_json = json.dumps(sorted_results, indent=4, cls=CustomJSONEncoder)
                
                with open(output_file, 'w') as json_file:
                    json_file.write(result_json)

    except yaml.YAMLError as exc:
        print(exc)