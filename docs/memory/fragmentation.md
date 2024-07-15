# Fragmentation in Memory Management 🚧

Fragmentation is a phenomenon in memory management where memory space is wasted and cannot be efficiently utilized. It occurs due to the allocation and deallocation of memory blocks in various sizes, leading to inefficient use of available memory. Fragmentation is categorized into two main types: internal and external fragmentation. Let’s dive deep into each type, their causes, and strategies to mitigate them.

#### 1. Internal Fragmentation 🧱

**Definition**: Internal fragmentation occurs when the memory allocated to a process is slightly larger than the memory required by the process. The excess allocated space within a memory block is wasted.

- **Causes**: Internal fragmentation typically happens in systems using fixed-size memory allocation methods, such as paging or fixed partitioning.
- **Example**: If a process requires 18 KB of memory and the system allocates memory in fixed blocks of 20 KB, then 2 KB of memory is wasted within the allocated block.

**Implications**:

- **Wasted Memory**: The unused memory within allocated blocks adds up, leading to significant memory wastage.
- **Reduced Efficiency**: Even though there might be enough total free memory, processes might still face allocation issues due to wastage within allocated blocks.

**Mitigation Strategies**:

- **Dynamic Allocation**: Allocating memory dynamically based on process requirements rather than fixed blocks can reduce internal fragmentation.
- **Buddy System**: Allocates memory in power-of-two sizes and allows merging adjacent free blocks, reducing internal fragmentation.

#### 2. External Fragmentation 🚧

**Definition**: External fragmentation occurs when free memory is scattered in small blocks between allocated memory segments, making it difficult to allocate contiguous memory blocks to processes.

- **Causes**: External fragmentation happens in systems using variable-sized memory allocation methods, such as dynamic partitioning.
- **Example**: If a system has 3 blocks of free memory (10 KB, 20 KB, and 15 KB) scattered across the memory space, a process requiring 30 KB of contiguous memory cannot be allocated, even though there is a total of 45 KB free.

**Implications**:

- **Memory Utilization**: External fragmentation reduces the efficiency of memory utilization, as scattered free blocks may be insufficient for large memory requests.
- **Allocation Failures**: Processes requiring large contiguous memory blocks may fail to allocate memory, despite having sufficient total free memory.

**Mitigation Strategies**:

- **Compaction 📏**: Moving allocated memory blocks to create larger contiguous free memory blocks. Compaction is a time-consuming process and usually performed during system idle times.
- **Paging 📜**: Dividing memory into fixed-sized pages eliminates external fragmentation, as any free page can be used to satisfy memory requests.
- **Segmentation ✂️**: Dividing memory into variable-sized segments based on logical divisions helps in managing memory more efficiently and reducing external fragmentation.
- **Buddy System 👯**: Combining both splitting and coalescing of memory blocks helps to maintain contiguous free memory and reduce fragmentation.

#### 3. Advanced Fragmentation Concepts 🔍

- **Memory Pools 🏊**: Allocating fixed-size memory blocks from pre-defined pools based on common request sizes. This reduces both internal and external fragmentation by fitting requests to suitable-sized blocks.
- **Garbage Collection 🚮**: Automatic reclamation of unused memory can help in reducing fragmentation. It involves identifying and collecting memory that is no longer in use and making it available for future allocations.
- **Slab Allocation 🧱**: Used primarily in kernel memory allocation, slab allocation manages caches of objects of the same size, reducing internal fragmentation and overhead.

#### 4. Performance Considerations 🚀

- **Overhead**: Techniques to mitigate fragmentation (like compaction) can introduce overhead and impact system performance. Efficient algorithms are crucial to balance the trade-off between reducing fragmentation and maintaining performance.
- **Real-time Systems**: In real-time systems, fragmentation can cause unpredictable memory allocation times, impacting system predictability and performance. Strategies like fixed-size allocation and real-time garbage collection are essential in such environments.
- **Embedded Systems**: Fragmentation management is critical in embedded systems with limited memory resources. Efficient allocation strategies and minimizing fragmentation are vital for optimal system performance.

# External Fragmentation in Memory Management 🚧

External fragmentation is a phenomenon in memory management where free memory is scattered in small, non-contiguous blocks, making it difficult to allocate large contiguous blocks of memory to processes, even if there is enough total free memory. This leads to inefficient memory utilization and can cause performance degradation.

#### Key Concepts and Causes 📘

1. **Definition**:

   - **External Fragmentation**: Occurs when free memory is divided into small, dispersed blocks due to the allocation and deallocation of memory blocks of varying sizes.
   - **Example**: Suppose there are three free memory blocks of 10 KB, 20 KB, and 30 KB, and a process requests a contiguous block of 50 KB. Even though the total free memory is 60 KB, the request cannot be fulfilled because there is no single contiguous block of 50 KB.
2. **Causes**:

   - **Variable-Sized Allocation**: Allocating memory blocks of different sizes leads to gaps between allocated blocks.
   - **Frequent Allocation and Deallocation**: Continuous allocation and deallocation of memory blocks create small gaps of free memory.
   - **Memory Deallocation Order**: The order in which memory blocks are freed can lead to fragmentation if contiguous blocks are not freed simultaneously.

#### Visualization of External Fragmentation 🔍

Consider the following example where memory is allocated and deallocated over time:

1. **Initial Allocation**:
   - Total Memory: 100 KB
   - Process A: 20 KB
   - Process B: 30 KB
   - Process C: 10 KB
   - Free Memory: 40 KB

```
|---- A ----|---- B ----|---- C ----|---- Free ----|
```

2. **Deallocate Process B**:
   - Process B is deallocated, creating a free block of 30 KB.
   - Free Memory: 30 KB + 40 KB = 70 KB

```
|---- A ----|---- Free (30 KB) ----|---- C ----|---- Free (40 KB) ----|
```

3. **Allocate Process D (25 KB)**:
   - Process D is allocated in the free block of 30 KB.
   - Free Memory: 5 KB + 40 KB = 45 KB

```
|---- A ----|---- D ----|---- Free (5 KB) ----|---- C ----|---- Free (40 KB) ----|
```

4. **Deallocate Process A and Process C**:
   - Processes A and C are deallocated, creating fragmented free memory blocks.
   - Free Memory: 20 KB + 5 KB + 10 KB + 40 KB = 75 KB

```
|---- Free (20 KB) ----|---- D ----|---- Free (5 KB) ----|---- Free (10 KB) ----|---- Free (40 KB) ----|
```

5. **Allocate Process E (50 KB)**:
   - Process E requires a contiguous block of 50 KB, but the largest contiguous free block is only 40 KB.
   - Allocation fails due to external fragmentation, despite having a total of 75 KB free.

#### Consequences of External Fragmentation ⚠️

1. **Wasted Memory**:

   - Free memory is wasted because it cannot be used to satisfy large memory requests.
2. **Reduced System Performance**:

   - Frequent allocation failures and the need for memory compaction can degrade system performance.
3. **Inefficient Memory Utilization**:

   - Memory is not utilized optimally, leading to potential underutilization of available resources.

### Solving External Fragmentation in Memory Management 🚧

External fragmentation occurs when free memory is divided into small, non-contiguous blocks, making it difficult to allocate large contiguous blocks of memory. Here are several techniques to solve or mitigate external fragmentation, along with a code example for one of the methods.

#### Techniques to Solve External Fragmentation 🔄

1. **Compaction 📏**:

   - **Definition**: Compaction involves moving allocated memory blocks to create larger contiguous blocks of free memory.
   - **Drawbacks**: It can be time-consuming and may cause performance overhead due to the need to move data around in memory.

   ```python
   class MemoryBlock:
       def __init__(self, start, size, allocated=False):
           self.start = start  # Starting address of the memory block
           self.size = size  # Size of the memory block
           self.allocated = allocated  # Allocation status of the block

   class MemoryManager:
       def __init__(self, total_memory):
           # Initialize memory manager with a single free block of total_memory size
           self.memory = [MemoryBlock(0, total_memory)]  

       def allocate(self, size):
           # Iterate over memory blocks to find a suitable free block
           for block in self.memory:
               if not block.allocated and block.size >= size:
                   if block.size > size:
                       # If the block is larger than needed, split it into allocated and free parts
                       new_block = MemoryBlock(block.start + size, block.size - size)
                       self.memory.insert(self.memory.index(block) + 1, new_block)
                   block.size = size  # Set the block size to the requested size
                   block.allocated = True  # Mark the block as allocated
                   return block.start  # Return the starting address of the allocated block
           raise MemoryError("No sufficient free memory available")  # Raise error if no suitable block is found

       def free(self, address):
           # Iterate over memory blocks to find the block with the given starting address
           for block in self.memory:
               if block.start == address:
                   block.allocated = False  # Mark the block as free
                   self.compact()  # Compact the memory to merge adjacent free blocks
                   return
           raise ValueError("Invalid address")  # Raise error if no block with the given address is found

       def compact(self):
           free_memory = 0  # Initialize total free memory counter
           allocated_blocks = []  # List to hold allocated blocks

           # Iterate over memory blocks to collect allocated blocks and count total free memory
           for block in self.memory:
               if block.allocated:
                   allocated_blocks.append(block)
               else:
                   free_memory += block.size

           # Clear current memory layout
           self.memory = []
           current_address = 0  # Initialize current address counter

           # Reallocate collected allocated blocks to new memory layout
           for block in allocated_blocks:
               block.start = current_address  # Update block's starting address
               current_address += block.size  # Increment current address by block size
               self.memory.append(block)

           # Add a single free block with the total free memory at the end of the memory layout
           if free_memory > 0:
               self.memory.append(MemoryBlock(current_address, free_memory))

   # Example usage
   if __name__ == "__main__":
       manager = MemoryManager(1024)  # Initialize memory manager with 1024 units of memory

       # Allocate memory blocks
       addr1 = manager.allocate(100)
       print(f"Allocated 100 units at address {addr1}")

       addr2 = manager.allocate(200)
       print(f"Allocated 200 units at address {addr2}")

       addr3 = manager.allocate(50)
       print(f"Allocated 50 units at address {addr3}")

       # Free a memory block and compact
       manager.free(addr1)
       print(f"Freed 100 units from address {addr1}")

       # Allocate another block to see the effect of compaction
       addr4 = manager.allocate(150)
       print(f"Allocated 150 units at address {addr4}")
   ```
2. **Paging 📜**:

   - **Definition**: Paging divides memory into fixed-size pages, eliminating the need for contiguous memory allocation.
   - **How It Works**: Memory is divided into pages and frames, and any page can be placed in any frame.
   - **Advantage**: Completely eliminates external fragmentation.

   ```python
   class PageTableEntry:
       def __init__(self, frame_number=None, valid=False):
           self.frame_number = frame_number  # The frame number this page maps to
           self.valid = valid  # Indicates whether this page table entry is valid

   class PagingMemoryManager:
       def __init__(self, page_size, num_frames):
           self.page_size = page_size  # Size of each page in memory
           self.num_frames = num_frames  # Total number of frames in physical memory
           self.page_table = {}  # Page table mapping virtual pages to physical frames
           self.frames = [None] * num_frames  # Physical frames, initially all are free (None)

       def allocate(self, num_pages):
           frame_indices = []  # List to store the indices of allocated frames
           for _ in range(num_pages):
               try:
                   # Find the first free frame (None) in the frames list
                   frame_index = self.frames.index(None)
                   self.frames[frame_index] = True  # Mark the frame as allocated
                   frame_indices.append(frame_index)
               except ValueError:
                   raise MemoryError("No free frames available")  # Raise an error if no free frames are available

           page_indices = list(range(len(self.page_table), len(self.page_table) + num_pages))  # List of new virtual page indices
           for page_index, frame_index in zip(page_indices, frame_indices):
               # Create a page table entry for each new virtual page, mapping it to the allocated frame
               self.page_table[page_index] = PageTableEntry(frame_number=frame_index, valid=True)
           return page_indices  # Return the list of allocated virtual page indices

       def free(self, page_indices):
           for page_index in page_indices:
               # Get the frame number for the given virtual page index
               frame_index = self.page_table[page_index].frame_number
               self.frames[frame_index] = None  # Mark the frame as free
               del self.page_table[page_index]  # Remove the page table entry

       def translate(self, page_index, offset):
           entry = self.page_table.get(page_index)  # Get the page table entry for the given page index
           if not entry or not entry.valid:
               raise MemoryError("Invalid page access")  # Raise an error if the page is not valid
           return entry.frame_number * self.page_size + offset  # Calculate and return the physical address

   # Example usage
   if __name__ == "__main__":
       manager = PagingMemoryManager(page_size=4096, num_frames=16)  # Initialize with 4 KB pages and 16 frames

       # Allocate memory blocks
       pages1 = manager.allocate(2)
       print(f"Allocated pages {pages1}")

       pages2 = manager.allocate(3)
       print(f"Allocated pages {pages2}")

       # Translate a virtual address to a physical address
       physical_address = manager.translate(pages1[0], 200)
       print(f"Translated virtual address (page {pages1[0]}, offset 200) to physical address {physical_address}")

       # Free memory blocks
       manager.free(pages1)
       print(f"Freed pages {pages1}")
   ```

   This version of the code includes detailed comments that explain each part of the process, from initializing page table entries and the paging memory manager to allocating and freeing memory, as well as translating virtual addresses to physical addresses.
3. **Segmentation ✂️**:

   - **Definition**: Segmentation divides memory into segments based on logical divisions, such as code, data, and stack.
   - **How It Works**: Each segment is allocated independently, reducing the impact of external fragmentation.
   - **Advantage**: Provides a logical view of memory, which can be more intuitive for programmers.

   ```python
   class Segment:
       def __init__(self, base, limit):
           self.base = base  # Starting address of the segment
           self.limit = limit  # Size of the segment

   class SegmentationMemoryManager:
       def __init__(self, total_memory):
           self.total_memory = total_memory  # Total memory available for allocation
           self.segments = {}  # Dictionary to hold allocated segments
           self.next_free_address = 0  # Next free address for allocation

       def allocate(self, size):
           if self.next_free_address + size > self.total_memory:
               raise MemoryError("No sufficient free memory available")  # Check if there is enough free memory
           segment = Segment(base=self.next_free_address, limit=size)  # Create a new segment
           segment_id = len(self.segments)  # Generate a unique segment ID
           self.segments[segment_id] = segment  # Store the segment in the segments dictionary
           self.next_free_address += size  # Update the next free address
           return segment_id  # Return the segment ID

       def free(self, segment_id):
           if segment_id in self.segments:
               del self.segments[segment_id]  # Remove the segment from the dictionary
           else:
               raise ValueError("Invalid segment ID")  # Raise an error if the segment ID is invalid

       def translate(self, segment_id, offset):
           segment = self.segments.get(segment_id)  # Get the segment for the given ID
           if not segment or offset >= segment.limit:
               raise MemoryError("Invalid segment access")  # Raise an error if the segment is invalid or offset is out of bounds
           return segment.base + offset  # Calculate and return the physical address

   # Example usage
   if __name__ == "__main__":
       manager = SegmentationMemoryManager(total_memory=1024)  # Initialize with 1024 units of memory

       # Allocate memory segments
       seg1 = manager.allocate(100)
       print(f"Allocated segment {seg1} with size 100")

       seg2 = manager.allocate(200)
       print(f"Allocated segment {seg2} with size 200")

       # Translate a logical address to a physical address
       physical_address = manager.translate(seg1, 50)
       print(f"Translated logical address (segment {seg1}, offset 50) to physical address {physical_address}")

       # Free memory segments
       manager.free(seg1)
       print(f"Freed segment {seg1}")
   ```
4. **Buddy System 👯**:

   - **Definition**: Allocates memory in power-of-two sizes and combines or splits memory blocks as needed.
   - **How It Works**: When a block of memory is freed, it is merged with its "buddy" (another free block of the same size) if possible.
   - **Advantage**: Reduces fragmentation by merging adjacent free blocks.

   ```python
   class BuddyAllocator:
       def __init__(self, size):
           self.size = size  # Total size of the memory managed by the allocator
           self.free_lists = {size: [(0, size)]}  # Free lists keyed by block size, initially containing the whole memory

       def allocate(self, size):
           size = self._next_power_of_two(size)  # Round up the requested size to the next power of two
           if size > self.size:
               raise MemoryError("Requested size is too large")  # Raise an error if the requested size exceeds total memory

           for block_size in sorted(self.free_lists.keys()):  # Iterate over sorted block sizes in free lists
               if block_size >= size and self.free_lists[block_size]:  # Find a free block that fits the requested size
                   block = self.free_lists[block_size].pop(0)  # Remove the block from the free list
                   self._split_block(block, block_size, size)  # Split the block if necessary
                   return block[0]  # Return the starting address of the allocated block
           raise MemoryError("No sufficient free memory available")  # Raise an error if no suitable block is found

       def free(self, address, size):
           size = self._next_power_of_two(size)  # Round up the size to the next power of two
           block = (address, size)  # Create a block tuple with address and size
           while True:
               buddy_address = self._find_buddy(block[0], block[1])  # Find the buddy address of the block
               buddy = (buddy_address, block[1])  # Create a buddy block tuple
               if buddy in self.free_lists.get(block[1], []):  # Check if the buddy is free
                   self.free_lists[block[1]].remove(buddy)  # Remove the buddy from the free list
                   block = (min(block[0], buddy[0]), block[1] * 2)  # Merge the block with its buddy
               else:
                   break  # Exit the loop if the buddy is not free
           if block[1] not in self.free_lists:
               self.free_lists[block[1]] = []  # Initialize the free list for the block size if not already present
           self.free_lists[block[1]].append(block)  # Add the merged block to the free list

       def _next_power_of_two(self, x):
           return 1 << (x - 1).bit_length()  # Calculate the next power of two greater than or equal to x

       def _split_block(self, block, block_size, size):
           while block_size > size:
               block_size //= 2  # Halve the block size
               buddy = (block[0] + block_size, block_size)  # Create a buddy block
               if block_size not in self.free_lists:
                   self.free_lists[block_size] = []  # Initialize the free list for the block size if not already present
               self.free_lists[block_size].append(buddy)  # Add the buddy to the free list

       def _find_buddy(self, address, size):
           return address ^ size  # Calculate the buddy address using the XOR operation

   # Example usage
   if __name__ == "__main__":
       allocator = BuddyAllocator(1024)  # Initialize a buddy allocator with 1024 units of memory

       # Allocate memory blocks
       addr1 = allocator.allocate(100)
       print(f"Allocated 100 units at address {addr1}")

       addr2 = allocator.allocate(200)
       print(f"Allocated 200 units at address {addr2}")

       addr3 = allocator.allocate(50)
       print(f"Allocated 50 units at address {addr3}")

       # Free memory blocks
       allocator.free(addr1, 100)
       print(f"Freed 100 units from address {addr1}")

       allocator.free(addr3, 50)
       print(f"Freed 50 units from address {addr3}")

       # Allocate another block to see the effect of merging
       addr4 = allocator.allocate(150)
       print(f"Allocated 150 units at address {addr4}")
   ```
5. **Memory Pools 🏊**:

   - **Definition**: Pre-allocates fixed-size blocks of memory for specific types of data.
   - **How It Works**: Different memory pools are used for different types of allocations.
   - **Advantage**: Reduces fragmentation by grouping similar-sized allocations together.

   ```python
   class MemoryPool:
       def __init__(self, block_size, num_blocks):
           self.block_size = block_size  # Size of each memory block
           self.num_blocks = num_blocks  # Total number of memory blocks
           self.free_blocks = list(range(num_blocks))  # List of indices of free blocks
           self.pool = [None] * num_blocks  # Pool to hold the memory blocks (for illustrative purposes)

       def allocate(self):
           if not self.free_blocks:
               raise MemoryError("No free blocks available")  # Raise an error if no free blocks are available
           block_index = self.free_blocks.pop(0)  # Get the index of the first free block
           return block_index  # Return the index of the allocated block

       def free(self, block_index):
           if block_index < 0 or block_index >= self.num_blocks:
               raise ValueError("Invalid block index")  # Raise an error if the block index is invalid
           self.free_blocks.append(block_index)  # Add the block index back to the list of free blocks

   # Example usage
   if __name__ == "__main__":
       pool = MemoryPool(block_size=64, num_blocks=10)  # Initialize a memory pool with 64-byte blocks and 10 blocks

       # Allocate memory blocks
       block1 = pool.allocate()
       print(f"Allocated block {block1}")

       block2 = pool.allocate()
       print(f"Allocated block {block2}")

       # Free a memory block
       pool.free(block1)
       print(f"Freed block {block1}")

       # Allocate another block to see the effect of reusing freed block
       block3 = pool.allocate()
       print(f"Allocated block {block3}")
   ```

# Internal Fragmentation in Memory Management 🧩

#### Key Concepts and Causes 📘

1. **Definition**:

   - **Internal Fragmentation**: The wasted space within an allocated memory block that arises because the block is larger than the requested memory. It occurs when memory is allocated in fixed-size blocks, and the allocated block is not fully utilized.
   - **Example**: If a process requests 22 KB of memory and is allocated a 32 KB block, the 10 KB of unused space within that block represents internal fragmentation.
2. **Causes**:

   - **Fixed-Size Allocation**: When memory is divided into fixed-size blocks or pages, and the memory request does not perfectly fit into these blocks, leading to leftover space.
   - **Paging**: In a paging system, memory is divided into pages (e.g., 4 KB each). If a process does not use the entire page, the unused portion results in internal fragmentation.
   - **Buddy System**: Allocates memory in power-of-two sizes. If the requested size is not a power of two, the allocated block will be the next power of two, causing internal fragmentation.

#### Consequences of Internal Fragmentation ⚠️

1. **Wasted Memory**:

   - **Inefficiency**: The unused space within allocated blocks results in inefficient memory utilization. Over time, this can add up to a significant amount of wasted memory.
   - **Example**: If many processes have similar inefficiencies, the cumulative wasted memory can be substantial.
2. **Reduced System Performance**:

   - **Paging and Swapping**: Wasted memory reduces the total available memory for other processes, potentially leading to increased paging or swapping.
   - **Example**: Frequent paging and swapping due to insufficient available memory can degrade overall system performance.

#### Techniques to Mitigate Internal Fragmentation 🔄

1. **Variable-Size Allocation**:

   - **Concept**: Allocates memory blocks that closely match the requested size, reducing the amount of unused space.
   - **Example**: Memory allocation techniques like best-fit, first-fit, and next-fit try to allocate memory blocks that are as close in size to the requested memory as possible.
2. **Sub-Page Allocation**:

   - **Concept**: Divides pages into smaller blocks to reduce internal fragmentation.
   - **Example**: Allocating memory in sub-page units (e.g., 512 bytes) within a 4 KB page to better match the requested sizes.
3. **Heap Management Algorithms**:

   - **Concept**: Use sophisticated algorithms to manage dynamic memory allocation and reduce fragmentation.
   - **Example**: Advanced heap management algorithms like slab allocation, which is used for managing memory for frequently used objects of different sizes, can help minimize internal fragmentation.

#### Examples of Techniques to Mitigate Internal Fragmentation 🌟

1. **Variable-Size Allocation**:

   - **Best-Fit Allocation**: Allocates the smallest block that is large enough for the request, minimizing wasted space.
   - **First-Fit Allocation**: Allocates the first block that is large enough, which may not always minimize fragmentation but is faster.
   - **Next-Fit Allocation**: Similar to first-fit but starts searching from the last allocated block, potentially distributing fragmentation more evenly.
2. **Sub-Page Allocation**:

   - **Concept**: Within each page, memory can be further divided into smaller units to match the size of the requests more closely.
   - **Example**: Instead of allocating whole pages to small objects, a sub-page allocator divides each page into smaller fixed-size blocks.
3. **Slab Allocation**:

   - **Concept**: Memory is divided into caches, each of which stores objects of a specific size.
   - **Example**: Slab allocation is used in kernel memory allocation where objects of the same type are grouped into slabs. Each slab contains multiple objects, and the slab allocator efficiently manages free and allocated objects.



Certainly! Here are the code implementations for various methods to solve internal fragmentation, each with meaningful comments explaining the process.

### 1. Variable-Size Allocation

**Best-Fit Allocation**:
Allocates the smallest block that is large enough for the request.

```python
class MemoryBlock:
    def __init__(self, start, size):
        self.start = start
        self.size = size
        self.free = True

class BestFitAllocator:
    def __init__(self, total_memory):
        self.total_memory = total_memory
        self.memory_blocks = [MemoryBlock(0, total_memory)]

    def allocate(self, size):
        best_fit = None
        for block in self.memory_blocks:
            if block.free and block.size >= size:
                if best_fit is None or block.size < best_fit.size:
                    best_fit = block
      
        if best_fit is None:
            raise MemoryError("No sufficient free memory available")
      
        if best_fit.size > size:
            new_block = MemoryBlock(best_fit.start + size, best_fit.size - size)
            self.memory_blocks.insert(self.memory_blocks.index(best_fit) + 1, new_block)
        best_fit.size = size
        best_fit.free = False
        return best_fit.start

    def free(self, address):
        for block in self.memory_blocks:
            if block.start == address:
                block.free = True
                self._merge_free_blocks()
                return
        raise ValueError("Invalid address")

    def _merge_free_blocks(self):
        merged_blocks = []
        previous_block = None
        for block in self.memory_blocks:
            if previous_block and previous_block.free and block.free:
                previous_block.size += block.size
            else:
                if previous_block:
                    merged_blocks.append(previous_block)
                previous_block = block
        if previous_block:
            merged_blocks.append(previous_block)
        self.memory_blocks = merged_blocks

# Example usage
if __name__ == "__main__":
    allocator = BestFitAllocator(1024)  # Initialize memory allocator with 1024 units of memory

    # Allocate memory blocks
    addr1 = allocator.allocate(100)
    print(f"Allocated 100 units at address {addr1}")

    addr2 = allocator.allocate(200)
    print(f"Allocated 200 units at address {addr2}")

    addr3 = allocator.allocate(50)
    print(f"Allocated 50 units at address {addr3}")

    # Free a memory block
    allocator.free(addr1)
    print(f"Freed 100 units from address {addr1}")

    # Allocate another block to see the effect of best-fit allocation
    addr4 = allocator.allocate(150)
    print(f"Allocated 150 units at address {addr4}")
```

### 2. Sub-Page Allocation

**Sub-Page Allocation**:
Divides pages into smaller blocks to reduce internal fragmentation.

```python
class SubPageAllocator:
    def __init__(self, page_size, sub_page_size, num_pages):
        self.page_size = page_size
        self.sub_page_size = sub_page_size
        self.num_pages = num_pages
        self.pages = [[None] * (page_size // sub_page_size) for _ in range(num_pages)]

    def allocate(self, size):
        num_sub_pages = (size + self.sub_page_size - 1) // self.sub_page_size  # Round up to nearest sub-page count
        for page_index, page in enumerate(self.pages):
            free_count = 0
            for sub_page_index, sub_page in enumerate(page):
                if sub_page is None:
                    free_count += 1
                    if free_count == num_sub_pages:
                        start_index = sub_page_index - free_count + 1
                        for i in range(start_index, start_index + num_sub_pages):
                            page[i] = True
                        return (page_index, start_index)
                else:
                    free_count = 0
        raise MemoryError("No sufficient free memory available")

    def free(self, page_index, start_index, size):
        num_sub_pages = (size + self.sub_page_size - 1) // self.sub_page_size  # Round up to nearest sub-page count
        for i in range(start_index, start_index + num_sub_pages):
            self.pages[page_index][i] = None

# Example usage
if __name__ == "__main__":
    allocator = SubPageAllocator(page_size=4096, sub_page_size=512, num_pages=16)  # Initialize with 4 KB pages and 512-byte sub-pages

    # Allocate memory blocks
    addr1 = allocator.allocate(1000)
    print(f"Allocated 1000 bytes at page {addr1[0]}, sub-page {addr1[1]}")

    addr2 = allocator.allocate(1500)
    print(f"Allocated 1500 bytes at page {addr2[0]}, sub-page {addr2[1]}")

    # Free a memory block
    allocator.free(addr1[0], addr1[1], 1000)
    print(f"Freed 1000 bytes from page {addr1[0]}, sub-page {addr1[1]}")

    # Allocate another block to see the effect of sub-page allocation
    addr3 = allocator.allocate(500)
    print(f"Allocated 500 bytes at page {addr3[0]}, sub-page {addr3[1]}")
```

### 3. Heap Management Algorithms

**Heap Management with First-Fit Allocation**:
Uses the first-fit strategy to allocate memory blocks.

```python
class MemoryBlock:
    def __init__(self, start, size):
        self.start = start
        self.size = size
        self.free = True

class FirstFitAllocator:
    def __init__(self, total_memory):
        self.total_memory = total_memory
        self.memory_blocks = [MemoryBlock(0, total_memory)]

    def allocate(self, size):
        for block in self.memory_blocks:
            if block.free and block.size >= size:
                if block.size > size:
                    new_block = MemoryBlock(block.start + size, block.size - size)
                    self.memory_blocks.insert(self.memory_blocks.index(block) + 1, new_block)
                block.size = size
                block.free = False
                return block.start
        raise MemoryError("No sufficient free memory available")

    def free(self, address):
        for block in self.memory_blocks:
            if block.start == address:
                block.free = True
                self._merge_free_blocks()
                return
        raise ValueError("Invalid address")

    def _merge_free_blocks(self):
        merged_blocks = []
        previous_block = None
        for block in self.memory_blocks:
            if previous_block and previous_block.free and block.free:
                previous_block.size += block.size
            else:
                if previous_block:
                    merged_blocks.append(previous_block)
                previous_block = block
        if previous_block:
            merged_blocks.append(previous_block)
        self.memory_blocks = merged_blocks

# Example usage
if __name__ == "__main__":
    allocator = FirstFitAllocator(1024)  # Initialize memory allocator with 1024 units of memory

    # Allocate memory blocks
    addr1 = allocator.allocate(100)
    print(f"Allocated 100 units at address {addr1}")

    addr2 = allocator.allocate(200)
    print(f"Allocated 200 units at address {addr2}")

    addr3 = allocator.allocate(50)
    print(f"Allocated 50 units at address {addr3}")

    # Free a memory block
    allocator.free(addr1)
    print(f"Freed 100 units from address {addr1}")

    # Allocate another block to see the effect of first-fit allocation
    addr4 = allocator.allocate(150)
    print(f"Allocated 150 units at address {addr4}")
```

### 4. Slab Allocation

**Slab Allocation**:
Divides memory into caches, each storing objects of a specific size.

```python
class Slab:
    def __init__(self, object_size, num_objects):
        self.object_size = object_size
        self.num_objects = num_objects
        self.free_objects = list(range(num_objects))
        self.slab = [None] * num_objects

class SlabAllocator:
    def __init__(self):
        self.caches = {}

    def create_cache(self, object_size, num_objects):
        self.caches[object_size] = Slab(object_size, num_objects)

    def allocate(self, object_size):
        if object_size not in self.caches:
            raise MemoryError("No cache available for the requested object size")
        slab = self.caches[object_size]
        if not slab.free_objects:
            raise MemoryError("No free objects available in the cache")
        object_index = slab.free_objects.pop(0)
        slab.slab[object_index] = True
        return object_index

    def free(self, object_size, object_index):
        if object_size not in self.caches or object_index >= self.caches[object_size].num_objects:
            raise ValueError("Invalid object size or index")
        slab = self.caches[object_size]
        slab.slab[object_index] = None
        slab.free_objects.append(object_index)

# Example usage
if __name__ == "__main__":
    allocator = SlabAllocator()
    allocator.create_cache(64, 10)  # Create a cache for 64-byte objects with 10 objects
    allocator.create_cache(128, 5)  # Create a cache for 128-byte objects with 5 objects

    # Allocate objects from the caches
    obj1 = allocator.allocate(64)
    print(f"Allocated 64-byte object at index {obj1}")

    obj2 = allocator.allocate(128)
    print(f"Allocated 

128-byte object at index {obj2}")

    # Free objects back to the caches
    allocator.free(64, obj1)
    print(f"Freed 64-byte object at index {obj1}")

    allocator.free(128, obj2)
    print(f"Freed 128-byte object at index {obj2}")

    # Allocate another object to see the effect of slab allocation
    obj3 = allocator.allocate(64)
    print(f"Allocated 64-byte object at index {obj3}")
```

### Summary

Here are the implementations for various methods to solve internal fragmentation:

1. **Variable-Size Allocation (Best-Fit)**: Allocates the smallest block that is large enough for the request.
2. **Sub-Page Allocation**: Divides pages into smaller blocks to better match the size of the requests.
3. **Heap Management Algorithms (First-Fit)**: Uses the first-fit strategy to allocate memory blocks.
4. **Slab Allocation**: Divides memory into caches, each storing objects of a specific size, to manage memory more efficiently.
