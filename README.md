# Zip File Extractor
Unzip file into folder of the same name and delete .zip after extraction

# Build:
git clone https://github.com/friday963/zip_file_extractor.git  
cd zip_file_extractor/  
go build  

# Usage
Place the executable in a folder with .zip files and execute it.  It will extract the file into a folder with the same name and dump the contents inside it. After it iterates over every .zip file it will delete the .zip files leaving you with just the folder and the contents inside it.
