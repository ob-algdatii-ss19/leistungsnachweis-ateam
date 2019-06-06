# GitHub-Page for documentation

GitHub deploys the content of this branch ```gh-pages``` automatically
to the GitHub-Page: https://ob-algdatii-ss19.github.io/leistungsnachweis-ateam/

## Usage

With each push to the branch ```gh-pages``` GitHub automatically generates a new GitHub-Page.

### Configuration

The configuration like title or theme can be done in the file ```_config.yml```

### Create a sub-page

To create a subpage you need to create a new markdown file.
At the top of file you need to add:

```
---
layout: default
title: [title of the]
description: [description is optionally]
---
```

To switch between the pages you can easily add the standard markdown links.
The file ending needs to be replaced with ```.html```.

### Styling

For styling your page, e.g. create tables or insert images you can have
a look at https://guides.github.com/features/mastering-markdown/.