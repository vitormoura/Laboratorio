Virtual Shared Folder
==================================

Web api to store and retrive files. Just a sample web project using GOLang.


Installing as a Windows Service
-------------------------------

1. Install and Build your project following the instructions of GOXC in https://github.com/laher/goxc

2. Download the NSSM on http://nssm.cc/

3. Execute the following commands: 

```Bash
nssm install VirtualSF virtualsf.exe 
nssm set VirtualSF Description "Serviço de publicação e consulta de arquivos compartilhados para aplicações web"
```


Uninstall Windows Service
-------------------------------


1. Execute the following nssm command:

```Bash
nssm remove VirtualSF
```