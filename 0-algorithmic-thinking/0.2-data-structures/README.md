# Data Structures

## Random Access Memory (RAM)

Variables are stored in random access memory (RAM). We sometimes call RAM "working memory" or just "memory".

In addition to "memory," your computer has storage (sometimes called "persistent storage" or "disk"). Storage is where we keep files like mp3s, videos, Word documents, and even executable programs or apps.

Memory (or RAM) is faster but has less space, while storage (or "disk") is slower but has more space. A modern laptop might have ~500GB of storage but only ~16GB of RAM.

![Processor - Memory controller - RAM](assets/processor-memory-controller-ram.svg)

Think of RAM like a really tall bookcase with a lot of shelves. Like, billions of shelves. 

The shelves are numbered. We call a shelf's number its address. Each shelf has one byte (8 bits) of storage.

We also have a processor that does all the real work inside our computer. It's connected to a memory controller. 

The memory controller does the actual reading and writing to and from RAM. 

Even though the memory controller can jump between far-apart memory addresses quickly, programs tend to access memory that's nearby. So computers are tuned to get an extra speed boost when reading memory addresses that're close to each other by using cache which even faster than RAM. 

When the processor asks for the contents of a given memory address, the memory controller also sends the contents of a handful of nearby memory addresses. And the processor puts all of it in the cache.