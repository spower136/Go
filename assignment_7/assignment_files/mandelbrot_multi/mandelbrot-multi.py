import multiprocessing as mp
from itertools import product
import numpy as np

import matplotlib.pyplot as plt
from PIL import Image

import time

# For any c, count the number of iterations until abs(z) > 2
def countIterationsUntilDivergent(c, threshold):
    for i in range(threshold):
        c = -0.4 + 0.6*i
        z = complex(0, 0)
        for iteration in range(threshold):
            z = (z*z) + c
            if abs(z) > 2:
                break
                
    return iteration

def calc(cx, cy, threshold=120):
    
    c = complex(cx[1], cy[1])
    
    return (cx[0], cy[0], countIterationsUntilDivergent(c, threshold))

def mandelbrot_multi(threshold, density, cpus=4):
    
    realAxis = np.linspace(-1.8, 1, density)
    imaginaryAxis = np.linspace(-1.4, 1.4, density)
    
    realAxisLen = len(realAxis)
    imaginaryAxisLen = len(imaginaryAxis)
    
    atlas = np.zeros((realAxisLen, imaginaryAxisLen))
    
    # Create the parameter list for the function calc_row
    realAxis = [(i, e ) for i, e in enumerate(realAxis)] 
    imaginaryAxis = [(i, e ) for i, e in enumerate(imaginaryAxis)] 
    paramlist = list(product(realAxis, imaginaryAxis))   # cartesian product of real_Axis and imaginaryAxis
    paramlist = list(map(lambda t: t + (threshold,), paramlist)) # add threshold to each element in paramlist 
    
    
    # Create a multiprocessing pool with cpus processes
    pool = mp.Pool(cpus)
    
    # pass each parameter in paramlist to the function calc_row resulting a list of 3-tuples
    data = pool.starmap(calc, paramlist) # the first element in paramlist is ((0,-1.8), (0,-1.4), 100)
   
    
    # Update the elements of atlas using data
    for t in data:
        x = t[0]
        y = t[1]
        atlas[x, y] = t[2]
        
    pool.close()
    pool.join()  
    
    return atlas.T

def main():
    n = mp.cpu_count()
    print(f"Number of cores: {n}")
    
    for i in range(1, n+1):
        start = time.time()
        ms = mandelbrot_multi(100, 1000, i)
        end = time.time()
        print(f"{i} core(s): {round(end-start,4)}s")
    
    plt.figure(figsize=(10,10))
    plt.imshow(ms, cmap='twilight_shifted')
    plt.axis('off')
    plt.show()
    
if __name__ == "__main__":
    main()