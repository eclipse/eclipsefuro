---
title: "Release1.20"
date: 2020-11-16T13:21:59+01:00
---

# 1.20.0
We should have done the release notes before, so this note will contain a summary of the older releases too.
In the fututre the messages will get smaller then this one.


### semver: 1.20.0
  date: 2020-11-13T10:12:52+01:00

#### changes:
   

- 9bf9e23970a3b68d969fa45b5c886a1469edc539

      note: 'feat: genBundledServiceProto to generate all specified services in one file installed services included'
      breaking: false
      description: genBundledServiceProto to generate all specified services in one file installed services included
- 5f2ec401c3e52df09efbb5c444db5e7a274d80b0

      note: 'feat: genBundledServiceProto to generate all specified services in one file'
      breaking: false
      description: genBundledServiceProto to generate all specified services in one file
 

### semver: 1.19.0
  date: 2020-11-05T11:38:11+01:00

#### changes:
    - 9bf9e23970a3b68d969fa45b5c886a1469edc539

      note: 'feat: genBundledServiceProto to generate all specified services in one file installed services included'
      breaking: false
      description: genBundledServiceProto to generate all specified services in one file installed services included

- 5f2ec401c3e52df09efbb5c444db5e7a274d80b0

      note: 'feat: genBundledServiceProto to generate all specified services in one file'
      breaking: false
      description: genBundledServiceProto to generate all specified services in one file
### semver: 1.18.3
  date: 2020-11-04T16:43:48+01:00

#### changes:
    - 0987f156be86b273b9d868148e273bc47eb4a9a2

      note: 'fix: no emitempty on yaml specs'
      breaking: false
      description: no emitempty on yaml specs
### semver: 1.18.2
  date: 2020-11-04T11:46:46+01:00

#### changes:
    - d6608ebb97c71996e846eeca9d09b9b530b19a4d

      note: 'fix: store google.type.Any as map[string]interface{} because it will not marshal to yaml else'
      breaking: false
      description: store google.type.Any as map[string]interface{} because it will not marshal to yaml else
### semver: 1.18.1
  date: 2020-11-03T09:58:53+01:00

#### changes:
    - 1f6cfe9f603b944e89c536161ea0c4cd4b10d6f3

      note: 'fix: ensure .furo file is copied even when it is not in the list'
      breaking: false
      description: ensure .furo file is copied even when it is not in the list
### semver: 1.18.0
  date: 2020-11-02T16:14:29+01:00

#### changes:
    - bdf5bb95c2ea838b7bc38b034299f3e77995fd6d

      note: 'feat: dist file management'
      breaking: false
      description: dist file management
### semver: 1.17.0
  date: 2020-11-02T09:50:52+01:00

#### changes:
    - 5993e41b5083bde5b99838222268bda1e978be72

      note: set version to 1.17.0
      breaking: false
      description: set version to 1.17.0
- b04c56fe93db81cf0cbfd53eae3eed3bfff523c4

      note: |-
      feat: escape newline, tabs and quotes in es6module so you can write the defaults like:

      default: |
      {
      "link": {
      "rel": "list",
      "href": "/api/xxx",
      "method": "GET",
      "type": "pkg.Type",
      "service": "Servicename"
      }
      }
      breaking: false
      description: 'escape newline, tabs and quotes in es6module so you can write the defaults like:'
      body: |-
      default: |
      {
      "link": {
      "rel": "list",
      "href": "/api/xxx",
      "method": "GET",
      "type": "pkg.Type",
      "service": "Servicename"
      }
      }
    - 85c22477a21538e71927eedeb7cd6997c0603f28

      note: 'fix: /* eslint-disable */ on esModule'
      breaking: false
      description: /* eslint-disable */ on esModule
### semver: 1.16.1
  date: 2020-10-30T13:10:02+01:00

#### changes:
    - e88ad37ad85139727db7957e83079e3e13412d12

      note: 'fix: version'
      breaking: false
      description: version
### semver: 1.16.0
  date: 2020-10-29T11:04:40+01:00

#### changes:
    - 605c2762a6c9c5fd457c1cdc9b226d67b76245f8

      note: |-
      feat: semver check against .furo config.
      A config which requires a newer release then installed, forces furo to quit.
      breaking: false
      description: semver check against .furo config.
### semver: 1.15.0
  date: 2020-10-28T16:50:47+01:00

#### changes:
    - eb1d7248817830bd461ab29c6250a4cda4caf910

      note: 'feat: imports are sorted and distinct'
      breaking: false
      description: imports are sorted and distinct
- bde2d0679a46362e870653670b6bfd917bc5e2d8

      note: 'feat: imports are sorted and distinct'
      breaking: false
      description: imports are sorted and distinct
### semver: 1.14.1
  date: 2020-10-25T14:57:21+01:00

#### changes:
    - 4d3178b7f4d62bba93af972f5097cd8ffdfed5ff

      note: 'fix: do not overwrite default values from spec when comming from muSpec'
      breaking: false
      description: do not overwrite default values from spec when comming from muSpec
- eb8be6794f6751336b406a9787cbf7972c085510

      note: 'fix: do not overwrite default values from spec when comming from muSpec'
      breaking: false
      description: do not overwrite default values from spec when comming from muSpec
- ddaffa49f8dcb59cf3f5b745a58f2014b8fb6a64

      note: 'docs: Documentation refined'
      
      conventional_commit:
      category: docs
      scope: ""
      breaking: false
      description: Documentation refined
- c669e95404fa37ba8fea93ba2984265e8c671d7a

      note: 'docs: sample files updated'
      
      conventional_commit:
      category: docs
      scope: ""
      breaking: false
      description: sample files updated
### semver: 1.14.0
  date: 2020-10-23T06:21:14+02:00

#### changes:
    - d01fdfdb0346ee06c7a54615f303a58ed5186bcb

      note: 'feat: additional path info in exported yaml'
      breaking: false
      description: additional path info in exported yaml
- cba77a277872d001a6d53967fabe736e80598dff

      note: 'feat: additional path info in exported yaml'
      breaking: false
      description: additional path info in exported yaml
### semver: 1.13.1
  date: 2020-10-19T10:22:16+02:00

#### changes:
    - b2cc22bcafe23008d20a8362afe1bb63bb05dd4d

      note: 'fix: version tag'
      breaking: false
      description: version tag
### semver: 1.13.0
  date: 2020-10-19T10:16:00+02:00

#### changes:
    - 4f7be62325741c58a5be96d0e9f175d2b404d9e8

      note: |-
      feat: using argument for flow, so you can write
      furo run flowname instead furo run -f=flowname
      breaking: false
      description: using argument for flow, so you can write
### semver: 1.12.1
  date: 2020-10-18T08:36:39+02:00

#### changes:
    - f6fcc8745a6a5c48cf39ce852a418e0ade2ef2b5

      note: 'fix: default field label with "package.type.field.label"'
      breaking: false
      description: default field label with "package.type.field.label"
- f1d2d971ed0f291499aa593a17afe9f2147ab9f2

      note: 'fix: default field label with "package.type.field.label"'
      breaking: false
      description: default field label with "package.type.field.label"
### semver: 1.12.0
  date: 2020-10-16T11:53:13+02:00

#### changes:
    - 7f9a41cb1c52c1b2db2e21684bd7e2034072052b

      note: 'feat: java package prefix "muSpec.javaPackagePrefix"'
      breaking: false
      description: java package prefix "muSpec.javaPackagePrefix"
### semver: 1.11.2
  date: 2020-10-16T07:52:38+02:00

#### changes:
    - ee63ec29b75e210330bdb7050d0263707e9f3aea

      note: 'fix: order of required, readonly when creating muSpec from spec'
      breaking: false
      description: order of required, readonly when creating muSpec from spec
- 127a0f212e862ee83480e3bb609f3561366a2458

      note: 'fix: order of required, readonly'
      breaking: false
      description: order of required, readonly
- 22b86c62c9a5b2302c7c23d41f191b1033ef5912

      note: 'fix: catch empty AstField.Constraints'
      breaking: false
      description: catch empty AstField.Constraints
### semver: 1.11.1
  date: 2020-10-15T11:32:39+02:00

#### changes:
    - 90ba2bbe61cb82586134d8c773ebd4fd74d24f99

      note: 'fix: catch empty __proto'
      breaking: false
      description: catch empty __proto
### semver: 1.11.0
  date: 2020-10-04T10:55:02+02:00

#### changes:
    - 529b8bf4f21b4e141da31b30cde916371ce9ee4c

      note: 'fix: removed the cleanbuild on protos, decision is not for the tools'
      breaking: false
      description: removed the cleanbuild on protos, decision is not for the tools
- bf61805029cfbc4476c9fe824b35cdcb76e07fcd

      note: 'feat: additional bindings on Update Request with method PUT and field update_mask'
      breaking: false
      description: additional bindings on Update Request with method PUT and field update_mask
### semver: 1.10.0
  date: 2020-09-25T15:28:48+02:00

#### changes:
    - eba42e0077fac5461ef21a8f958ce70dc99e73ff

      note: 'feat: type extensions'
      breaking: false
      description: type extensions
- 485c417c037c6756ddb58328aee90a9a1df7f4ba

      note: 'feat: service extensions'
      breaking: false
      description: service extensions
- d784ed39a556290b21ea73bc940eb7e797f651b3

      note: query description removed from es6spec
      breaking: false
      description: query description removed from es6spec
- 76e97fca57ce1622ff967e416255bdb78245d8f9

      note: 'chore: export installed types and services too'
      
      conventional_commit:
      category: chore
      scope: ""
      breaking: false
      description: export installed types and services too
- 70663701da57c79dd4153a64276f1436df98da59

      note: 'fix: version number'
      breaking: false
      description: version number
- a0080a3a389121c4f265ed69b2bd13d85938c601

      note: 'fix: depth=1 removed'
      breaking: false
      description: depth=1 removed
- 293a0171bcc7b3a7616fc02fc5738408acc3d8ed

      note: 'feat: use git command only, go-git removed'
      breaking: false
      description: use git command only, go-git removed
- dbe381db5cd6780cda46bbe949b3d941fb500807

      note: 'feat: use git command only, go-git removed'
      breaking: false
      description: use git command only, go-git removed
- 269fde54f0474ee2a19f939327fa192f9d5dc3e6

      note: 'feat: use git command only, go-git removed'
      breaking: false
      description: use git command only, go-git removed
- 86b5cde11d82ef1e9b84c4a266698477b6dc969a

      note: 'fix: git command fallback'
      breaking: false
      description: git command fallback
- 104284a5d71efba93786dc50ff97d381b972f97f

      note: 'fix: git command fallback'
      breaking: false
      description: git command fallback
- b2cf1734e3de672f4fa76df5d09bd2d5e9ba9fbe

      note: git command fallback
      breaking: false
      description: git command fallback
- 5a784513d2a89cb147076c9eb483de12e3c4e2fd

      note: git command fallback
      breaking: false
      description: git command fallback
- 8571a24af0797ff6cfdee0ae065e6c73dce94669

      note: feat bodyField can be named like you want. Default is "body"
      breaking: false
      description: feat bodyField can be named like you want. Default is "body"
- c73060005695d36bf8d5e42608013cad8348e4f9

      note: feat bodyField can be named like you want. Default is "body"
      breaking: false
      description: feat bodyField can be named like you want. Default is "body"
- 40d46c49d7a411d3a5ab1c9624cfca9955376f46

      note: feat muSrvSanitize => service sanitizer
      breaking: false
      description: feat muSrvSanitize => service sanitizer
- bb8b2d8d18b231e7d5ce7e46104beb2f776890fd

      note: |-
      feat better error message with sourcode pointer
      more robust regex
      breaking: false
      description: feat better error message with sourcode pointer
      body: more robust regex
    - 666579f481e2c10cde1a75cc4ae129b615c6d30a

      note: feat Spec2MuSpec creating of services
      breaking: false
      description: feat Spec2MuSpec creating of services
- a02c2b571a819cc73c75a33373aa8fd15cb3a162

      note: feat Spec2MuSpec creating of types
      breaking: false
      description: feat Spec2MuSpec creating of types
- adb1eb93bb01d15d8b882bd6f0f81af15c2c772a

      note: feat Spec2MuSpec link the objects
      breaking: false
      description: feat Spec2MuSpec link the objects
- b07412c1d86bced3393e618f8edf016ff246e10c

      note: feat Spec2MuSpec init
      breaking: false
      description: feat Spec2MuSpec init
- cc65d91559de89d5bd7290db55ec15d57a80bcad

      note: feat MuSpec2Spec doc update
      breaking: false
      description: feat MuSpec2Spec doc update
- c8f6e7c29d69193db433dbcd8cc947815a98656d

      note: feat MuSpec2Spec doc update
      breaking: false
      description: feat MuSpec2Spec doc update
- 35781541763e7ed2b2cbc84a0e2deab0b6cedde6

      note: feat MuSpec2Spec
      breaking: false
      description: feat MuSpec2Spec
- bd4efc6b8dc1165c79f819aa86461ef0ce1a827e

      note: fix checkimports with single root specs and protos
      breaking: false
      description: fix checkimports with single root specs and protos
- 8b9c529e1436edd03d23e6fe6e6f4dc92b3cfffb

      note: feat install -fresh || -f  to refresh the local repo (or force a fresh installation)
      breaking: false
      description: feat install -fresh || -f  to refresh the local repo (or force a fresh installation)
- 227e0ac06419d5494f6d029587897ef139854d42

      note: feat .notation cleanup for service request and response
      breaking: false
      description: feat .notation cleanup for service request and response
- 2fab65593a4b6d261085202e1d95cf54b93e71ec

      note: feat .notation cleanup for service request and response
      breaking: false
      description: feat .notation cleanup for service request and response
- ac1b05dbe0779f4f16dae35c874a34649e0b0ed6

      note: feat c++ type notation and resolution
      breaking: false
      description: feat c++ type notation and resolution
- c9d457b0a1cd0bf7cad06a6117aad3e1ab723e0d

      note: |-
      generate service protos without request messages
      init of ClassicServiceTemplate
      breaking: false
      description: generate service protos without request messages
      body: init of ClassicServiceTemplate
    - b0a4bd96d3ee92547f13303ce2fccff46e16debb

      note: |-
      generate protos without installed as default
      yaml exporter without installed as default
      breaking: false
      description: generate protos without installed as default
      body: yaml exporter without installed as default
    - 9d699b41aec986a51d481fc82fe294f9c4e479c9

      note: |-
      generate protos without installed as default
      yaml exporter without installed as default
      breaking: false
      description: generate protos without installed as default
      body: yaml exporter without installed as default
    - 563fcc28612c006a9dc34f096d0d1fe54ce8c464

      note: |-
      checkImports.go
      improved samples
      breaking: false
      description: checkImports.go
      body: improved samples
    - 90c7f58dcca2d06fef871b9b4dfed47070c9ee20

      note: small refactoring
      breaking: false
      description: small refactoring
- e4de2c7eeb9a187056d83803d1d2e2c4f4702569

      note: 'feat: init'
      breaking: false
      description: init
### semver: 1.9.4
  date: 2020-09-25T13:36:35+02:00

#### changes:
    - 485c417c037c6756ddb58328aee90a9a1df7f4ba

      note: 'feat: service extensions'
      breaking: false
      description: service extensions
### semver: 1.9.3
  date: 2020-09-24T14:33:19+02:00

#### changes:
    - d784ed39a556290b21ea73bc940eb7e797f651b3

      note: query description removed from es6spec
      breaking: false
      description: query description removed from es6spec
- 76e97fca57ce1622ff967e416255bdb78245d8f9

      note: 'chore: export installed types and services too'
      
      conventional_commit:
      category: chore
      scope: ""
      breaking: false
      description: export installed types and services too
- 70663701da57c79dd4153a64276f1436df98da59

      note: 'fix: version number'
      breaking: false
      description: version number
- a0080a3a389121c4f265ed69b2bd13d85938c601

      note: 'fix: depth=1 removed'
      breaking: false
      description: depth=1 removed
### semver: 1.9.2
  date: 2020-09-21T12:10:08+02:00

#### changes:
    - 293a0171bcc7b3a7616fc02fc5738408acc3d8ed

      note: 'feat: use git command only, go-git removed'
      breaking: false
      description: use git command only, go-git removed
- dbe381db5cd6780cda46bbe949b3d941fb500807

      note: 'feat: use git command only, go-git removed'
      breaking: false
      description: use git command only, go-git removed
- 269fde54f0474ee2a19f939327fa192f9d5dc3e6

      note: 'feat: use git command only, go-git removed'
      breaking: false
      description: use git command only, go-git removed
- 86b5cde11d82ef1e9b84c4a266698477b6dc969a

      note: 'fix: git command fallback'
      breaking: false
      description: git command fallback
### semver: 1.9.1
  date: 2020-09-21T11:19:03+02:00

#### changes:
    - 104284a5d71efba93786dc50ff97d381b972f97f

      note: 'fix: git command fallback'
      breaking: false
      description: git command fallback
- b2cf1734e3de672f4fa76df5d09bd2d5e9ba9fbe

      note: git command fallback
      breaking: false
      description: git command fallback
- 5a784513d2a89cb147076c9eb483de12e3c4e2fd

      note: git command fallback
      breaking: false
      description: git command fallback
### semver: 1.9.0
  date: 2020-09-20T07:03:08+02:00

#### changes:
    - 8571a24af0797ff6cfdee0ae065e6c73dce94669

      note: feat bodyField can be named like you want. Default is "body"
      breaking: false
      description: feat bodyField can be named like you want. Default is "body"
- c73060005695d36bf8d5e42608013cad8348e4f9

      note: feat bodyField can be named like you want. Default is "body"
      breaking: false
      description: feat bodyField can be named like you want. Default is "body"
### semver: 1.8.0
  date: 2020-09-19T22:45:29+02:00

#### changes:
    - 40d46c49d7a411d3a5ab1c9624cfca9955376f46

      note: feat muSrvSanitize => service sanitizer
      breaking: false
      description: feat muSrvSanitize => service sanitizer
- bb8b2d8d18b231e7d5ce7e46104beb2f776890fd

      note: |-
      feat better error message with sourcode pointer
      more robust regex
      breaking: false
      description: feat better error message with sourcode pointer
      body: more robust regex
    - 666579f481e2c10cde1a75cc4ae129b615c6d30a

      note: feat Spec2MuSpec creating of services
      breaking: false
      description: feat Spec2MuSpec creating of services
### semver: 1.7.0
  date: 2020-09-19T09:07:02+02:00

#### changes:
    - a02c2b571a819cc73c75a33373aa8fd15cb3a162

      note: feat Spec2MuSpec creating of types
      breaking: false
      description: feat Spec2MuSpec creating of types
- adb1eb93bb01d15d8b882bd6f0f81af15c2c772a

      note: feat Spec2MuSpec link the objects
      breaking: false
      description: feat Spec2MuSpec link the objects
- b07412c1d86bced3393e618f8edf016ff246e10c

      note: feat Spec2MuSpec init
      breaking: false
      description: feat Spec2MuSpec init
### semver: 1.6.1
  date: 2020-09-17T13:55:33+02:00

#### changes:
    - cc65d91559de89d5bd7290db55ec15d57a80bcad

      note: feat MuSpec2Spec doc update
      breaking: false
      description: feat MuSpec2Spec doc update
### semver: 1.6.0
  date: 2020-09-17T13:48:12+02:00

#### changes:
    - c8f6e7c29d69193db433dbcd8cc947815a98656d

      note: feat MuSpec2Spec doc update
      breaking: false
      description: feat MuSpec2Spec doc update
- 35781541763e7ed2b2cbc84a0e2deab0b6cedde6

      note: feat MuSpec2Spec
      breaking: false
      description: feat MuSpec2Spec
### semver: 1.5.1
  date: 2020-09-16T23:05:09+02:00

#### changes:
    - bd4efc6b8dc1165c79f819aa86461ef0ce1a827e

      note: fix checkimports with single root specs and protos
      breaking: false
      description: fix checkimports with single root specs and protos
### semver: 1.5.0
  date: 2020-09-16T11:06:49+02:00

#### changes:
    - 8b9c529e1436edd03d23e6fe6e6f4dc92b3cfffb

      note: feat install -fresh || -f  to refresh the local repo (or force a fresh installation)
      breaking: false
      description: feat install -fresh || -f  to refresh the local repo (or force a fresh installation)
- 227e0ac06419d5494f6d029587897ef139854d42

      note: feat .notation cleanup for service request and response
      breaking: false
      description: feat .notation cleanup for service request and response
- 2fab65593a4b6d261085202e1d95cf54b93e71ec

      note: feat .notation cleanup for service request and response
      breaking: false
      description: feat .notation cleanup for service request and response
- ac1b05dbe0779f4f16dae35c874a34649e0b0ed6

      note: feat c++ type notation and resolution
      breaking: false
      description: feat c++ type notation and resolution
- c9d457b0a1cd0bf7cad06a6117aad3e1ab723e0d

      note: |-
      generate service protos without request messages
      init of ClassicServiceTemplate
      breaking: false
      description: generate service protos without request messages
      body: init of ClassicServiceTemplate
### semver: 1.4.0
  date: 2020-09-14T14:12:52+02:00

#### changes:
    - b0a4bd96d3ee92547f13303ce2fccff46e16debb

      note: |-
      generate protos without installed as default
      yaml exporter without installed as default
      breaking: false
      description: generate protos without installed as default
      body: yaml exporter without installed as default
### semver: 1.3.0
  date: 2020-09-09T07:39:41+02:00

#### changes:
    - 9d699b41aec986a51d481fc82fe294f9c4e479c9

      note: |-
      generate protos without installed as default
      yaml exporter without installed as default
      breaking: false
      description: generate protos without installed as default
      body: yaml exporter without installed as default
    - 563fcc28612c006a9dc34f096d0d1fe54ce8c464

      note: |-
      checkImports.go
      improved samples
      breaking: false
      description: checkImports.go
      body: improved samples
    - 90c7f58dcca2d06fef871b9b4dfed47070c9ee20

      note: small refactoring
      breaking: false
      description: small refactoring
- e4de2c7eeb9a187056d83803d1d2e2c4f4702569

      note: 'feat: init'
      breaking: false
      description: init
- 445a9486f0bdfd6e7ae23a922be3dca961642f09

      note: 'fix: lowercase in generated env specs'
      breaking: false
      description: lowercase in generated env specs
- f32790ad542bf4e66ea40f6e473db65a304404a9

      note: 'fix: autocreate of collections with correct entity types'
      breaking: false
      description: autocreate of collections with correct entity types
- 9744caef00902e666033c472fc2ebef87113b58e

      note: 'feat: Release version 1.1.0'
      breaking: false
      description: Release version 1.1.0
- b4b823f7148ed17d88bc99189db0141b4cdeda38

      note: 'feat: calling furo without arguments will run the default flow.'
      breaking: false
      description: calling furo without arguments will run the default flow.
- 173aad3d53ef3a219bf7600a721bd51c7e962682

      note: 'doc: doc updated'
      
      conventional_commit:
      category: doc
      scope: ""
      breaking: false
      description: doc updated
- be922360dd2dd465c1d286beb6ccbef0d6a4e396

      note: 'feat: env var with FST prefix'
      breaking: false
      description: env var with FST prefix
- 628a8d7497ce509a18ad43cb9515082ccf1aafa4

      note: installation manual
      breaking: false
      description: installation manual
- fa8479243df579acaee329019ee5ab25d22c8760

      note: remove gobuild.sh, it is not needed
      breaking: false
      description: remove gobuild.sh, it is not needed
- 8b7add3707e5568e68003ea482bbcc4e15e4ec19

      note: test feat gobuild init
      breaking: false
      description: test feat gobuild init
- de25f7b5e35a9cfd430595fc286378b553b8db0e

      note: test feat gobuild init
      breaking: false
      description: test feat gobuild init
- a60f9f2d44e54bd38c8c701a9c50827da6d6a4c0

      note: feat gobuild init
      breaking: false
      description: feat gobuild init
- 767107987ba3405c8574987a3c2596222e4c2855

      note: 'fix: root command prints help as default'
      breaking: false
      description: root command prints help as default
- 4b063c8db78ffc07d5fbf5298546966ed4405136

      note: 'feat: custom config'
      breaking: false
      description: custom config
- 0a70b54fb489e9c45ab3bdcfe53ce5bab8261504

      note: 'feat: custom config'
      breaking: false
      description: custom config
- 5e010716787a1e336d3122ac2facefd499d5e08f

      note: 'feat: custom config'
      breaking: false
      description: custom config
- d2569a57c12708df1161b930b989e0abcce01bfd

      note: dev notice removed
      breaking: false
      description: dev notice removed
- f2a518c8115b042ac92d52721c2f533655bb40eb

      note: 'feat: Sample for a command'
      breaking: false
      description: Sample for a command
- ff473474589154b74e99f1682cc2735110032f35

      note: 'feat: Sample for a command'
      breaking: false
      description: Sample for a command
- 1e43d13267a692f866caf2972d09704f26b332a1

      note: 'doc: doc updated'
      
      conventional_commit:
      category: doc
      scope: ""
      breaking: false
      description: doc updated
- 9cac4490645db4a97b0eb2534a7d64c6007546ec

      note: 'doc: doc init'
      
      conventional_commit:
      category: doc
      scope: ""
      breaking: false
      description: doc init
- 2a39737c5b4af3ecb8b94ee49f07030a1f0becd9

      note: 'doc: doc init'
      
      conventional_commit:
      category: doc
      scope: ""
      breaking: false
      description: doc init
- edb253e39405f57411a626ca53e7f5ba22ad9965

      note: 'feat: running commands from flow'
      breaking: false
      description: running commands from flow
- 0f703ae97c6908d725ce19629c809e5a473cdbca

      note: 'fix: pb in package name'
      breaking: false
      description: pb in package name
- 3620fabfc5811790faa7c0a9ee410ec3869198a2

      note: export full tree as yaml
      breaking: false
      description: export full tree as yaml
- 7c86a5e44c7d617f31f0b0ecc5ce87eda998103b

      note: export full tree as yaml
      breaking: false
      description: export full tree as yaml
- 2c9e7c0f7a81bf5766726d5fadb413a9e8d3009a

      note: generate full env es module
      breaking: false
      description: generate full env es module
- 97cb8eeb592a595db724b8fa4e7509ec8aac4322

      note: installer with repo
      breaking: false
      description: installer with repo
- 94c761da852d3ea6c41056b2f7cc664bcb6ed7ab

      note: installer init
      breaking: false
      description: installer init
- cf8f14b406d6221eed75348cf45cc9d0223b96f7

      note: installer init
      breaking: false
      description: installer init
- 05c90b80de1d0cae81d4959181decbf36a37dc3a

      note: load dependencies
      breaking: false
      description: load dependencies
- c293c5aea4b348b2f16934270a3b21c5452b29ab

      note: load dependencies
      breaking: false
      description: load dependencies
- 4011063b4049be22fd1cd0640e08dee7ac17f454

      note: generate type protos and migration from initializr
      breaking: false
      description: generate type protos and migration from initializr
- af165db5efee37bd62e4def8cb076d926497869b

      note: generate type protos and migration from initializr
      breaking: false
      description: generate type protos and migration from initializr
- cca9368e8d4399bdc608426c520daaa028e4b014

      note: file watcher init
      breaking: false
      description: file watcher init
- c08480d3e2a1226362af2a5f1bb5c2fdaf48ac3c

      note: flow runner
      breaking: false
      description: flow runner
- 22a263a2f4e2bc5f2c1511ecf91eface236c576c

      note: flow runner
      breaking: false
      description: flow runner
- 8afedbe12a68cc6e06b48cfb87a2f1ff988796a0

      note: muService2Spec
      breaking: false
      description: muService2Spec
- 47176b1e4e338adc35c31259523f1165f0ce4ee3

      note: initial commit
      breaking: false
      description: initial commit
- b5efebc81e4a14c8a7bd00cb20ffda23389091ee

      note: initial commit
      breaking: false
      description: initial commit
### semver: 1.2.0
  date: 2020-09-07T15:31:44+02:00

#### changes:
    - 563fcc28612c006a9dc34f096d0d1fe54ce8c464

      note: |-
      checkImports.go
      improved samples
      breaking: false
      description: checkImports.go
      body: improved samples
    - 90c7f58dcca2d06fef871b9b4dfed47070c9ee20

      note: small refactoring
      breaking: false
      description: small refactoring
- e4de2c7eeb9a187056d83803d1d2e2c4f4702569

      note: 'feat: init'
      breaking: false
      description: init
- 445a9486f0bdfd6e7ae23a922be3dca961642f09

      note: 'fix: lowercase in generated env specs'
      breaking: false
      description: lowercase in generated env specs
- f32790ad542bf4e66ea40f6e473db65a304404a9

      note: 'fix: autocreate of collections with correct entity types'
      breaking: false
      description: autocreate of collections with correct entity types
- 9744caef00902e666033c472fc2ebef87113b58e

      note: 'feat: Release version 1.1.0'
      breaking: false
      description: Release version 1.1.0
- b4b823f7148ed17d88bc99189db0141b4cdeda38

      note: 'feat: calling furo without arguments will run the default flow.'
      breaking: false
      description: calling furo without arguments will run the default flow.
- 173aad3d53ef3a219bf7600a721bd51c7e962682

      note: 'doc: doc updated'
      
      conventional_commit:
      category: doc
      scope: ""
      breaking: false
      description: doc updated
- be922360dd2dd465c1d286beb6ccbef0d6a4e396

      note: 'feat: env var with FST prefix'
      breaking: false
      description: env var with FST prefix
- 628a8d7497ce509a18ad43cb9515082ccf1aafa4

      note: installation manual
      breaking: false
      description: installation manual
- fa8479243df579acaee329019ee5ab25d22c8760

      note: remove gobuild.sh, it is not needed
      breaking: false
      description: remove gobuild.sh, it is not needed
- 8b7add3707e5568e68003ea482bbcc4e15e4ec19

      note: test feat gobuild init
      breaking: false
      description: test feat gobuild init
- de25f7b5e35a9cfd430595fc286378b553b8db0e

      note: test feat gobuild init
      breaking: false
      description: test feat gobuild init
- a60f9f2d44e54bd38c8c701a9c50827da6d6a4c0

      note: feat gobuild init
      breaking: false
      description: feat gobuild init
- 767107987ba3405c8574987a3c2596222e4c2855

      note: 'fix: root command prints help as default'
      breaking: false
      description: root command prints help as default
- 4b063c8db78ffc07d5fbf5298546966ed4405136

      note: 'feat: custom config'
      breaking: false
      description: custom config
- 0a70b54fb489e9c45ab3bdcfe53ce5bab8261504

      note: 'feat: custom config'
      breaking: false
      description: custom config
- 5e010716787a1e336d3122ac2facefd499d5e08f

      note: 'feat: custom config'
      breaking: false
      description: custom config
- d2569a57c12708df1161b930b989e0abcce01bfd

      note: dev notice removed
      breaking: false
      description: dev notice removed
- f2a518c8115b042ac92d52721c2f533655bb40eb

      note: 'feat: Sample for a command'
      breaking: false
      description: Sample for a command
- ff473474589154b74e99f1682cc2735110032f35

      note: 'feat: Sample for a command'
      breaking: false
      description: Sample for a command
- 1e43d13267a692f866caf2972d09704f26b332a1

      note: 'doc: doc updated'
      
      conventional_commit:
      category: doc
      scope: ""
      breaking: false
      description: doc updated
- 9cac4490645db4a97b0eb2534a7d64c6007546ec

      note: 'doc: doc init'
      
      conventional_commit:
      category: doc
      scope: ""
      breaking: false
      description: doc init
- 2a39737c5b4af3ecb8b94ee49f07030a1f0becd9

      note: 'doc: doc init'
      
      conventional_commit:
      category: doc
      scope: ""
      breaking: false
      description: doc init
- edb253e39405f57411a626ca53e7f5ba22ad9965

      note: 'feat: running commands from flow'
      breaking: false
      description: running commands from flow
- 0f703ae97c6908d725ce19629c809e5a473cdbca

      note: 'fix: pb in package name'
      breaking: false
      description: pb in package name
- 3620fabfc5811790faa7c0a9ee410ec3869198a2

      note: export full tree as yaml
      breaking: false
      description: export full tree as yaml
- 7c86a5e44c7d617f31f0b0ecc5ce87eda998103b

      note: export full tree as yaml
      breaking: false
      description: export full tree as yaml
- 2c9e7c0f7a81bf5766726d5fadb413a9e8d3009a

      note: generate full env es module
      breaking: false
      description: generate full env es module
- 97cb8eeb592a595db724b8fa4e7509ec8aac4322

      note: installer with repo
      breaking: false
      description: installer with repo
- 94c761da852d3ea6c41056b2f7cc664bcb6ed7ab

      note: installer init
      breaking: false
      description: installer init
- cf8f14b406d6221eed75348cf45cc9d0223b96f7

      note: installer init
      breaking: false
      description: installer init
- 05c90b80de1d0cae81d4959181decbf36a37dc3a

      note: load dependencies
      breaking: false
      description: load dependencies
- c293c5aea4b348b2f16934270a3b21c5452b29ab

      note: load dependencies
      breaking: false
      description: load dependencies
- 4011063b4049be22fd1cd0640e08dee7ac17f454

      note: generate type protos and migration from initializr
      breaking: false
      description: generate type protos and migration from initializr
- af165db5efee37bd62e4def8cb076d926497869b

      note: generate type protos and migration from initializr
      breaking: false
      description: generate type protos and migration from initializr
- cca9368e8d4399bdc608426c520daaa028e4b014

      note: file watcher init
      breaking: false
      description: file watcher init
- c08480d3e2a1226362af2a5f1bb5c2fdaf48ac3c

      note: flow runner
      breaking: false
      description: flow runner
- 22a263a2f4e2bc5f2c1511ecf91eface236c576c

      note: flow runner
      breaking: false
      description: flow runner
- 8afedbe12a68cc6e06b48cfb87a2f1ff988796a0

      note: muService2Spec
      breaking: false
      description: muService2Spec
- 47176b1e4e338adc35c31259523f1165f0ce4ee3

      note: initial commit
      breaking: false
      description: initial commit
- b5efebc81e4a14c8a7bd00cb20ffda23389091ee

      note: initial commit
      breaking: false
      description: initial commit
### semver: 1.1.2
  date: 2020-09-05T22:13:03+02:00

#### changes:
    - 445a9486f0bdfd6e7ae23a922be3dca961642f09

      note: 'fix: lowercase in generated env specs'
      breaking: false
      description: lowercase in generated env specs
### semver: 1.1.1
  date: 2020-09-05T08:03:42+02:00

#### changes:
    - f32790ad542bf4e66ea40f6e473db65a304404a9

      note: 'fix: autocreate of collections with correct entity types'
      breaking: false
      description: autocreate of collections with correct entity types
### semver: 1.1.0
  date: 2020-09-04T07:24:13+02:00

#### changes:
    - 9744caef00902e666033c472fc2ebef87113b58e

      note: 'feat: Release version 1.1.0'
      breaking: false
      description: Release version 1.1.0
- b4b823f7148ed17d88bc99189db0141b4cdeda38

      note: 'feat: calling furo without arguments will run the default flow.'
      breaking: false
      description: calling furo without arguments will run the default flow.
- 173aad3d53ef3a219bf7600a721bd51c7e962682

      note: 'doc: doc updated'
      
      conventional_commit:
      category: doc
      scope: ""
      breaking: false
      description: doc updated
- be922360dd2dd465c1d286beb6ccbef0d6a4e396

      note: 'feat: env var with FST prefix'
      breaking: false
      description: env var with FST prefix
### semver: 1.0.0
  date: 2020-09-03T22:36:17+02:00

#### changes:
    - 628a8d7497ce509a18ad43cb9515082ccf1aafa4

      note: installation manual
      breaking: false
      description: installation manual
- fa8479243df579acaee329019ee5ab25d22c8760

      note: remove gobuild.sh, it is not needed
      breaking: false
      description: remove gobuild.sh, it is not needed
- 8b7add3707e5568e68003ea482bbcc4e15e4ec19

      note: test feat gobuild init
      breaking: false
      description: test feat gobuild init
- de25f7b5e35a9cfd430595fc286378b553b8db0e

      note: test feat gobuild init
      breaking: false
      description: test feat gobuild init
- a60f9f2d44e54bd38c8c701a9c50827da6d6a4c0

      note: feat gobuild init
      breaking: false
      description: feat gobuild init
- 767107987ba3405c8574987a3c2596222e4c2855

      note: 'fix: root command prints help as default'
      breaking: false
      description: root command prints help as default
- 4b063c8db78ffc07d5fbf5298546966ed4405136

      note: 'feat: custom config'
      breaking: false
      description: custom config
- 0a70b54fb489e9c45ab3bdcfe53ce5bab8261504

      note: 'feat: custom config'
      breaking: false
      description: custom config
- 5e010716787a1e336d3122ac2facefd499d5e08f

      note: 'feat: custom config'
      breaking: false
      description: custom config
- d2569a57c12708df1161b930b989e0abcce01bfd

      note: dev notice removed
      breaking: false
      description: dev notice removed
- f2a518c8115b042ac92d52721c2f533655bb40eb

      note: 'feat: Sample for a command'
      breaking: false
      description: Sample for a command
- ff473474589154b74e99f1682cc2735110032f35

      note: 'feat: Sample for a command'
      breaking: false
      description: Sample for a command
- 1e43d13267a692f866caf2972d09704f26b332a1

      note: 'doc: doc updated'
      
      conventional_commit:
      category: doc
      scope: ""
      breaking: false
      description: doc updated
- 9cac4490645db4a97b0eb2534a7d64c6007546ec

      note: 'doc: doc init'
      
      conventional_commit:
      category: doc
      scope: ""
      breaking: false
      description: doc init
- 2a39737c5b4af3ecb8b94ee49f07030a1f0becd9

      note: 'doc: doc init'
      
      conventional_commit:
      category: doc
      scope: ""
      breaking: false
      description: doc init
- edb253e39405f57411a626ca53e7f5ba22ad9965

      note: 'feat: running commands from flow'
      breaking: false
      description: running commands from flow
- 0f703ae97c6908d725ce19629c809e5a473cdbca

      note: 'fix: pb in package name'
      breaking: false
      description: pb in package name
- 3620fabfc5811790faa7c0a9ee410ec3869198a2

      note: export full tree as yaml
      breaking: false
      description: export full tree as yaml
- 7c86a5e44c7d617f31f0b0ecc5ce87eda998103b

      note: export full tree as yaml
      breaking: false
      description: export full tree as yaml
- 2c9e7c0f7a81bf5766726d5fadb413a9e8d3009a

      note: generate full env es module
      breaking: false
      description: generate full env es module
- 97cb8eeb592a595db724b8fa4e7509ec8aac4322

      note: installer with repo
      breaking: false
      description: installer with repo
- 94c761da852d3ea6c41056b2f7cc664bcb6ed7ab

      note: installer init
      breaking: false
      description: installer init
- cf8f14b406d6221eed75348cf45cc9d0223b96f7

      note: installer init
      breaking: false
      description: installer init
- 05c90b80de1d0cae81d4959181decbf36a37dc3a

      note: load dependencies
      breaking: false
      description: load dependencies
- c293c5aea4b348b2f16934270a3b21c5452b29ab

      note: load dependencies
      breaking: false
      description: load dependencies
- 4011063b4049be22fd1cd0640e08dee7ac17f454

      note: generate type protos and migration from initializr
      breaking: false
      description: generate type protos and migration from initializr
- af165db5efee37bd62e4def8cb076d926497869b

      note: generate type protos and migration from initializr
      breaking: false
      description: generate type protos and migration from initializr
- cca9368e8d4399bdc608426c520daaa028e4b014

      note: file watcher init
      breaking: false
      description: file watcher init
- c08480d3e2a1226362af2a5f1bb5c2fdaf48ac3c

      note: flow runner
      breaking: false
      description: flow runner
- 22a263a2f4e2bc5f2c1511ecf91eface236c576c

      note: flow runner
      breaking: false
      description: flow runner
- 8afedbe12a68cc6e06b48cfb87a2f1ff988796a0

      note: muService2Spec
      breaking: false
      description: muService2Spec
