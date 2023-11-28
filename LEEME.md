# Un CLI de explainshell.com hecho con Go

Esto es un proyecto personal para uso propio. A veces necesitas conocer un comando o parámetro específico de la shell. Man is muy pesado y tldr demasiado ligero. explainshell.com va justo al grano.

## Algunos ejemplos

```code
$ expl tar zxvf example.tar.gz
The GNU version of the tar archiving utility
-z, --gzip, --gunzip --ungzip
-x, --extract, --get
      extract files from an archive
-v, --verbose
      verbosely list files processed
-f, --file ARCHIVE
      use archive file or device ARCHIVE
```

```code
$ expl du -sh
estimate file space usage
-s, --summarize
       display only a total for each argument
-h, --human-readable
       print sizes in human readable format (e.g., 1K 234M 2G)
```

## Instalar / Uso

No necesita instalación. Es sólo un ejecutable. Hay versiones compiladas en bin/ para linux, win64 y win32 (útil si estás usando Cygwin).

Para Linux copia el fichero expl dentro de tu $HOME/bin o cualquier directorio bin de tu $PATH

## Compilar por ti mismo

Golang debe de estar instalado.

Clonar el repositorio git y ejecutar el script compile.sh. 3 ficheros (expl, expl32.exe y expl.exe) se volverán a crear en bin/
