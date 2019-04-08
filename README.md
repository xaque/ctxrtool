# CTXR Image Converter
This is a WIP conversion tool for CTXR images.

I only know of the CTXR image format through reverse engineering, the actual format may be under another name and may even be well known. I just found these files with a .ctxr extension and wasn't able to open them with any image tools I had. But, I was able to figure out they were in a pretty simple image format.

I haven't figured out what all the header values are, so the tool currently just makes assumptions about values. Specifically, it assumes the image stores pixels in ARGB order and each sample is 1 byte. It also assumes the width and height attributes are at a hardcoded index in the header. This works for a few example files I had lying around but not everything. The example image file provided will convert propertly.

`convertCTXR.go` takes any number of arguments which are .ctxr files and converts them to .pam files

`ctxr_to_png.sh` runs the above then uses the `netpbm` package to convert the .pam file to .png
