---
title: "Moving to eclipse"
date: 2021-07-12T21:22:22+01:00
---

We started to move the many parts of the furo project from theNorstroem to eclipse last month. 

This forced us to restructure the individual parts of *furo*. 
The restructuring definitely brings us some advantages for the project itself. 

For the consumers, the sharp separation of the involved components,  a cleaner naming and improved documentation will 
give them a better view of the big picture.

We kept the changes that you have to made as low as possible. 
Please read in the guide below, what you have to do, to migrate your project.

## Migration guide: 

### Web related stuff:
There are no changes to make, all npm packages have the same name.

### API and spec related stuff:
All you have to do is 

- Rename the `.spectools` file to `.furo`. 
- Replace the word `spectools` with `furo` in the renamed file.
- Update the dependency to the `furoBaseSpecs` (if you have) to at least `git@github.com:theNorstroem/furoBaseSpecs.git v1.27.1`.
- If you work with BEC, install the latest or at least `1.27.3` by running `docker pull thenorstroem/furo-bec:v1.27.3` and you are done.
- If you work with a local installation, update the brew tap and run `brew install furo` and you are done.
- If you installed furo from the source, checkout `github.com/eclipse/eclipsefuro`.

{{< hint info>}}
**Note on nested spces**
If you have dependenies which are not already switched from spectools to furo, you have to rename the `.spectools` 
file to `.furo` in your dependency folder.   
{{< /hint >}}


### Custom generators
If you have built custom generators for `furoc`, 
change the dependencies for spectools and furoc to at least `github.com/eclipse/eclipsefuro v1.27.3`

	github.com/theNorstroem/furoc  becomes github.com/eclipse/eclipsefuro/furoc
	github.com/theNorstroem/spectools becomes github.com/eclipse/eclipsefuro/furo

