---
weight: 8
title: "install"
---

## furo install

Installs the configured dependencies from the .furo config to the specified folder.

Enter the name of the repository and a tag or branch to install. If you want latest (not recommended) add **main** (**master**). 

```yaml
dependencies:
  - "git@github.com:theNorstroem/furoBaseSpecs.git v1.27.1" # The importer looks for all **/*.type.spec files recursive The importer looks for all **/*.service.spec files recursive
dependenciesDir: dependencies # Your installed dependencies from other spec project are saved in this directory

```

It is up to you to add the `dependencies` folder to your version control system. 
We strongly recommend this if you work with packages that are not owned by 
your own organisation. So you will have them always available, even when the source vanishes. 



```
furo install [flags]
```

### Options

```
  -f, --fresh   rebuild the package directories
  -h, --help    help for install
```

### Options inherited from parent commands

```
      --config string   config file (default is CWD/.furo.yaml)
```


