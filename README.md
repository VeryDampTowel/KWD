# KWD
KoboWordlistDownloader - GoLang Project to download "My Words" from the Kobo e-Readers


> Only tested with Kobo Libra 2 but should work with other Kobo devices

## Usage
1. Download the tool for your specific OS.
2. Extract KWD from the downloaded file
3. Plug your Kobo device into your computer
4. If you don't see the .kobo directory then you need to show hidden files - (`CMD+Shift+.` on Mac or `View > Show > Hidden items` on Windows.)
5. [Optional] I would recommend copying the KoboReader.sqlite to somewhere else like your Desktop to avoid any possibility of corrupting the database on your eReader.
6. Run the tool! `./KWD /file/path/to/KoboReader.sqlite`

You should see a `Kobo_Wordlist.txt` file with the list of words in your current directory.
