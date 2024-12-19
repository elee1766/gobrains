require(httr)
require(data.table)
require(jsonlite)
api_configs <- function() {
  conf <- new.env()
  conf[["labs_endpoint"]] = "https://labs.api.listenbrainz.org/"
  return(conf)
}

find_similar <- function(mbid, algo = "session_based_days_9000_session_300_contribution_5_threshold_15_limit_50_skip_30") {
  conf <- api_configs()
  full_url <- paste0(conf[["labs_endpoint"]],"similar-recordings/json")
  resp <- httr::GET(
    full_url,
    accept_json(),
    query = list(
      recording_mbids = mbid,
      algorithm = algo
      )
    )
  
  if(http_error(resp)) {
    print("failed getting recs", resp)
    return(NULL)
  }
  
  parsed <- content(resp, as="parsed")
  dt <- do.call(rbind, lapply(parsed, as.data.table))
  dt[,count:=1]
  dt
  return(dt)
}


