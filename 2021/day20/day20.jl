# Advent of Code, Day 20
#
# Transform an image by successively replacing pixels with values
# looked up from a translation table, the index being the value
# of the 9 cells surrounding each pixel, converted from binary to
# decimal. Much complicated by the fact that the input data has a
# 1 in the first position of the translation table, meaning that
# empty areas are filled with 1s, which muck up the pixel count.
#
# AK, 20/12/2021


# Read data into list of rows
f = "sample.txt"
f = "input.txt"
d = split(strip(read(f, String)), "\n")

# First row is "algorithm", convert # to 1 and . to 0
code = map(x -> x == '#' ? 1 : 0, [c for c in d[1]])
#println(code)

# Remaining rows are the image, convert to matrix
rows = map(r -> [c == '#' ? 1 : 0 for c in r], d[3:end])
image = Matrix(transpose(reshape(vcat(rows...), (:, length(rows)))))


# Transform an image by iteratively applying translation function
function transform(img, iterations)

    for iter in 1:iterations
        
        # If the code starts with a 1, it causes the whole image to
        # flip on each second iteration. To undo this, we neeed to
        # fill the border with 0 on the first iteration, 1 on the
        # second iteration, and so on.
        if code[1] == 1 && iseven(iter)
            bfill = 1
        else
            bfill = 0
        end

        # Enhance the image and display number of on bits
        img = enhance(img, bfill)
        #display(img)
        println("Iteration ", iter, ": bits set = ", sum(img))

    end
end


# Enhance the image, by converting 3x3 areas around a pixel
# to 9-digit binary number, converting this to decimal, and
# then looking up replacement bit in the "algorithm" vector
function enhance(img, padWith)

    # Make a copy of the image with a 2-thick border
    # around it, with the value of the padding alternating
    # to get around the flashing image problem when there
    # is a 1 at the beginning of the translation code.
    padded = addPadding(img, 2, padWith)

    # Make an empty array for the output, one bigger
    # on each side then the original image
    nrows, ncols = size(img)
    out = zeros(Int, nrows+2, ncols+2)

    # Go though the original image, and replace pixels
    # based on algorithm lookup
    for r in 1:nrows+2
        for c in 1:ncols+2
            n = numAt(padded, r+1, c+1)
            out[r,c] = code[Int(n+1)]
        end
    end

    #Return the enhanced  image
    return out
    
end


# Return the "number" centered at the given coordinates, i.e., 
# take the 3x3 cells around that number, string them along
# row-wise, and convert that binary representation to a decimal
# number
function numAt(img, r, c)

    # Get the 3x3 bits around the given coordinates
    # and turn them into a vector
    bits = reshape(transpose(img[r-1:r+1,c-1:c+1]), 1, :)

    # Convert this binary sequence to a decimal number
    #return bin2dec(bits)
    n = 0
    for b in bits
        n = n * 2 + b
    end
    return n

end


# Add a given number of layer of padding around a matrix, 
# using the provide fill value (1/0)
function addPadding(img, n, padWith)

    # Add padding on top and bottom
    nr, nc = size(img)
    v = padWith == 1 ? ones(Int, n, nc) : zeros(Int, n, nc)
    img = vcat(v, img, v)

    # Add padding on sides
    nr, nc = size(img)
    v = padWith == 1 ? ones(Int, nr, n) : zeros(Int, nr, n)
    img = hcat(v, img, v)

    return img

end


# Perform the two transformations (Part 1 has 2 iterations, Part 2 has 50).
# For sample data, should be 35 and 3351, respectively
println("PART 1:")
transform(image, 2)

println("\nPART 2:")
transform(image, 50)
