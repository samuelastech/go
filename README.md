# Go

This is a place to learn GoLang fundamentals and concepts using Mardown in Jupyter Notebook. Inside `src` you will find
all my notes.

## Usage

We need to install Go's kernel to run it in Jupyter Notebook — [Juan Nathaniel](https://levelup.gitconnected.com/running-golang-on-jupyter-notebook-f7f9fba37812) was my reference.

Install it by using **Docker** as it is OS-agnostic:

```shell
docker run -it -p 8888:8888 -v $(pwd)/src:/src gopherdata/gophernotes:latest-ds
```

As Juan explained:

> The *latest-ds* tag tells docker to pull a version of gophernotes package already installed with popular data science libraries, such as, GoNum, GoLearn, and GoDa.

A URL will be prompted in your terminal, copy and paste it into your browser.