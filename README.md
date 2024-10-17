# 2llamas

## What are the expand files?
Take input:
* p = Pick Up
* d = Drop Off
* w = Weight
* r = Rate
* t = Temperature

expand takes such params (any number) and converts to a more readable
format. It also automatically adds to clipboard, so you can past it 
wherever needed.

##### OS specifics
expand is for mac
expand.exe is for windows
expand.go is the source code

### On Mac
./expand [params]

example:
```
./expand p1200 d1200 w2000 r0.50 t32
```
### On Windows
expand.exe [params]

example:
```
expand.exe p1200 d1200 w2000 r0.50 t32
```

## Before running a .sh file for the first time
You will need to add persmission to execute the files. execute the
following code:

chmod 755 [filename].sh
## Running a .sh file

./[file name] parameters

in brackets (don't type brackets when typing out the command), put in
file name. parameters will be the "p2000 d0600 w8000" and so on
