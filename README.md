# What is this?

This program will read in your iTunes Library.xml file and then walk through
each track in the file.  For each track, it will update the modification time
on the actual file to match the "Date Added" timestamp that is recorded in the
iTunes library database.

Over the years there are many normal actions and activities that can cause the
modification time of a file to be updated long after it was added to iTunes.
It's very easy for these timestamps to become out of sync and to have the file
modification times much newer than the actual date the track was added to your
iTunes library.

# Why would I care about that?

If you have a program which is working off the iTunes media directory and not
working through the iTunes API. For example, if you are using [roon] to play
your music and you have your iTunes set up as a library location it can
optionally use the file time to track when the file was added.

# Cool, I use roon and this sounds fantastic!  What do I do?

* If you are on Catalina and using the newer "Music" app instead of iTunes, you
  should first manually export your library to an XML file.

    * File > Library > Export Library...
    * Save the `Library.XML` file somewhere

* Download a binary for your OS from the [releases page] page here and then
  copy it into the same directory as your `Library.XML` file.

* Open a terminal window and navigate to that same directory

* Run `./itunes-file-mtimes`

* In roon, go to Settings > Library > Import Settings and set "Import Date
  Defaults to" to "File Modification Time"

* In roon, go to "Storage" and perform a "Force Rescan" on your attached iTunes
  library location.

# Useful Links

* [Discussion on Roon Community forum](https://community.roonlabs.com/t/fixed-my-itunes-library-date-added-to-work-better-with-roon-library-imports/93034?u=nugget)

# Credits

Written by David "nugget" McNett in Go.

[roon]: https://roonlabs.com
[releases page]: https://github.com/nugget/itunes-file-mtimes/releases
