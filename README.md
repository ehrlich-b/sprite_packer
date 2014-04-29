This program takes a directory and packs all of the images (.jif, .png, and .gif) into a sprite that can be used by css.

Free icon sets to test with: http://www.hongkiat.com/blog/free-icon-sets-to-bookmark/

Running the program:
- - -
    git clone https://github.com/adotout/sprite_packer.git
    cd sprite_packer
    go get "github.com/adotout/pack_2d"
    go run sprite_pack.go "[icon folder]"
You should now have a file in the current directory with a file called "sprite.png" with all of the icons packed together into a sprite.

TODO
- - -
* Output css style sheet
* Command line flags: output file name, directory search