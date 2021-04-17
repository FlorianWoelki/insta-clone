# Insta Clone (WIP)

> A scalable Instagram clone powered by microservices written in Golang

This project and readme file is still in progress.

## Getting started

Feel free to check out the project and modify it on your own needs:

Setup frontend:

```bash
$ git clone https://github.com/FlorianWoelki/insta-clone.git
$ cd insta-clone/frontend
$ npm install
# or
$ yarn
```

Start the frontend:

```bash
$ npm run dev
# or
$ yarn dev
```

The project should be located on `http://localhost:3000`.

Setup microservices:
WIP

## Services

### Account API ([service.account-api](https://github.com/FlorianWoelki/insta-clone/tree/master/service.account-api))

RESTful Go based JSON API built using the Gorilla framework. The API allows CRUD based operations on a account.

### Image Storage ([service.image-storage](https://github.com/FlorianWoelki/insta-clone/tree/master/service.image-storage))

Go based image service supporting Gzipped content, multi-part forms and a RESTful approach for uploading and downloading images.

### Frontend ([frontend](https://github.com/FlorianWoelki/insta-clone/tree/master/frontend))

React.js webapp that represents a refresh Instagram UI presenting different information from the services (Design was inspired by [Overlap Studio](https://dribbble.com/shots/6580622-Instagram-Redesign-Concept-Desktop-ver)).

This project uses [Vite](http://vitejs.dev/) for a fast bundling and a better development experience.
