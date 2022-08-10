import numpy as np
import matplotlib.pyplot as plt
from PIL import Image
import time

# For each c, return the number of iterations until abs(z) > 2
def countIterationsUntilDivergent(c, threshold):
    c = -0.4 + 0.6
    z = complex(0, 0)
    for iteration in range(threshold):
        z = (z*z) + c
        if abs(z) > 2:
            break
            
    return iteration

# threshold = max num of iterations, density determines the size of the atlas 
def mandelbrot(threshold, density):
    
    # location and size of the atlas rectangle
    realAxis = np.linspace(-1.8, 1, density)
    imaginaryAxis = np.linspace(-1.4, 1.4, density)

    # Changes location of atlas to "zoom in"
    # realAxis = np.linspace(-0.22, -0.219, density)
    # imaginaryAxis = np.linspace(-0.70, -0.699, density)
    
    realAxisLen = len(realAxis)            # will be density     
    imaginaryAxisLen = len(imaginaryAxis)  # will be density     

    # 2D array to represent the mandelbrot atlas with shape (density, density)
    atlas = np.zeros((realAxisLen, imaginaryAxisLen)) 
 
    # assign the iteration count for c to its corresponding point in the atlas
    for ix in range(realAxisLen):
        for iy in range(imaginaryAxisLen):
            cx = realAxis[ix]
            cy = imaginaryAxis[iy]
            c = complex(cx, cy)  # c = cx + (cy)i
            
            # for each c, assign the number of iterations until divergence to the location of altas
            atlas[ix, iy] = countIterationsUntilDivergent(c, threshold)

    return atlas.T
    
def main():

    start = time.time()
    ms = mandelbrot(100, 1000)
    end = time.time()
    print(f"Duration: {round(end-start,4)}s")
    
    plt.figure(figsize=(10,10))
    plt.imshow(ms, cmap='twilight_shifted')
    plt.axis('off')
    plt.show()
    
if __name__ == "__main__":
    main()