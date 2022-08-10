import numpy as np
import matplotlib.pyplot as plt
import time

def mandelbrot(h_range, w_range, theshold):
    
    # create two 2D arrays: y with the shape (h_range, 1) and x with the shape (1, w_range)
    y, x = np.ogrid[-1.4: 1.4: h_range*1j, -1.8: 1: w_range*1j]

    # Changes location of atlas to "zoom in"
    # y, x = np.ogrid[-0.7: -0.699: h_range*1j, -0.22: -0.219: w_range*1j]

    a_array = x + y*1j # a 2D array with shape (h_range, w_range)

    z_array = np.zeros(a_array.shape)
    atlas = theshold + np.zeros(a_array.shape) # set all values in atlas to threshold

    for i in range(theshold):
        # mandelbrot equation
        z_array = z_array**2 + a_array

        # make a boolean array for diverging indicies of z_array
        z_size_array = z_array * np.conj(z_array)
        divergent_array = z_size_array > 4  # assig True to all the locatons at which abs(z)>2
        
        # assign i to the all the locations at which the value in divergent_array is True
        atlas[divergent_array] = i  
        
        # prevent overflow for diverging locations 
        z_array[divergent_array] = 0 # but it makes the image look different from the others
    
    return atlas

def main():
    start = time.time()
    ms = mandelbrot(1000, 1000, 100)
    end = time.time()
    print(f"Duration: {round(end-start,4)}s")
    
    plt.figure(figsize=(10,10))
    plt.imshow(ms, cmap='twilight_shifted')
    plt.axis('off')
    plt.show()
    
if __name__ == "__main__":
    main()