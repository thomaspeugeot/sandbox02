# sandbox02
test bed for Ng compilation issue reprodution

## pre requisites

Angular 10.2
```

     _                      _                 ____ _     ___
    / \   _ __   __ _ _   _| | __ _ _ __     / ___| |   |_ _|
   / △ \ | '_ \ / _` | | | | |/ _` | '__|   | |   | |    | |
  / ___ \| | | | (_| | |_| | | (_| | |      | |___| |___ | |
 /_/   \_\_| |_|\__, |\__,_|_|\__,_|_|       \____|_____|___|
                |___/
    

Angular CLI: 10.2.0
Node: 10.22.1
OS: darwin x64

Angular: 
... 
Ivy Workspace: 

Package                      Version
------------------------------------------------------
@angular-devkit/architect    0.1002.0 (cli-only)
@angular-devkit/core         10.2.0 (cli-only)
@angular-devkit/schematics   10.2.0 (cli-only)
@schematics/angular          10.2.0 (cli-only)
@schematics/update           0.1002.0 (cli-only)
```


## build Ng librairies

### animah librairy
```
cd animah/ng
npm install
ng build animah
ng build animahcontrol
```

### gorgo library
```
cd gorgo/ng
npm install
ng build gorgo
ng build gorgodiagrams
```

### laundromat library
```
cd laundromat/ng
npm install
ng build laundromat
cd ../..
```

## application laundromat

```
cd laundromat/ng
ng serve
```

Unexpected behavior

Note : to have the go backend run
```
cd gorgo/go/models
go install
cd ../../..
```
```
cd laundromat/go/models
go install
cd ../../..
```

```
go run main.go -client-control
```


