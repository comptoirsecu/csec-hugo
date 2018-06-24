# csec-hugo

## How to work locally

Requirements : `docker`, `docker-compose`
Build site and start a webserver + gulp watch:

    git clone https://github.com/comptoirsecu/csec-hugo.git
    cd csec-hugo
    docker-compose up


Check `docker-compose.yml` and docker/ folder for more information.
Two container will be lauched: hugo server and yarn + gulp watch. No need to
install anything else on your computer, container are self sufficient.

If you need to use the hugo binary:

    docker-compose run --rm csec-hugo <command>

    # Example:
    docker-compose run --rm csec-hugo hugo version


If you need to use anything else (gulp, yarn...):

    docker-compose run --rm csec-dependencies <command>

    # Example:
    docker-compose run --rm csec-dependencies yarn update


The docker-compose.yml will mount current folder in containers at runtime.
