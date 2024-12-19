source("libs/data_getters.r")

start_seed <- "7db2c09d-345f-46e6-900e-0f4499e29c3d"
initial_pop <- find_similar(start_seed)

create_generations <- function(x) {
  generations <- list() 
  generations[[1]] = initial_pop[,.(recording_mbid,count)]
  samples <- 5
  ## three iterations total, two hops
  for(i in 2:3) {
    # sample from the set 
    current_sample <- generations[[i-1]][sample(.N, samples)]
    # generate more samples, keeping the first 5 samples in the list
    current_gen <- data.table::copy(current_sample)
    for(t in 1:nrow(current_sample)) {
      new <- find_similar(current_sample[1]$recording_mbid[[1]])
      new <- new[,.(recording_mbid,count)]
      current_gen <- rbind(current_gen, new)
    }
    current_gen <- current_gen[,.(count=sum(count)), by="recording_mbid"]
    generations[[i]] <- current_gen
  }
  return(generations)
}

created_generations <- lapply(1:100, create_generations)

source("libs/similarity.r")
final_playlists <- lapply(created_generations, \(x) x[[2]]$recording_mbid)
calculate_similarity_metrics(final_playlists)
