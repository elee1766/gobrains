# Required libraries
library(data.table)

# Function to calculate various similarity metrics between sets
calculate_similarity_metrics <- function(set_list) {
  n <- length(set_list)
  
  # Initialize results data.table
  results <- data.table(
    set_name = names(set_list),
    jaccard = numeric(n),
    dice = numeric(n),
    overlap = numeric(n),
    cosine = numeric(n)
  )
  
  # Calculate metrics for each set compared to all others combined
  for (i in 1:n) {
    # Get current set
    current_set <- set_list[[i]]
    
    # Combine all other sets
    other_sets_indices <- (1:n)[-i]
    all_other_elements <- unique(unlist(set_list[other_sets_indices]))
    
    # Calculate intersection and sizes
    intersection_size <- length(intersect(current_set, all_other_elements))
    union_size <- length(union(current_set, all_other_elements))
    
    # Calculate metrics
    results[i, `:=`(
      # Jaccard Index = |A ∩ B| / |A ∪ B|
      jaccard = intersection_size / union_size,
      
      # Dice Coefficient = 2|A ∩ B| / (|A| + |B|)
      dice = 2 * intersection_size / (length(current_set) + length(all_other_elements)),
      
      # Overlap Coefficient = |A ∩ B| / min(|A|, |B|)
      overlap = intersection_size / min(length(current_set), length(all_other_elements)),
      
      # Cosine Similarity = |A ∩ B| / sqrt(|A| * |B|)
      cosine = intersection_size / sqrt(length(current_set) * length(all_other_elements))
    )]
  }
  
  return(results[])  # Adding [] to ensure printing
}
