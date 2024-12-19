# gobrains

tl;dr - this is a project to try to make a iheart/pandora like radio, using metabrainz+spotify

## origin story

gobrains is code i am writing to interact with listenbrainz/musicbrainz, (and perhaps other metabrainz related sergices) using go.

originally i helped worked on a client for spotify, which cumulated in https://git.asdf.cafe/abs3nt/gspot

the key feature is the "autoradio" feature, which allows you infinitely refill a radio station with reccomendations based off some criterium, like pandora/iheart

as the spotify reccomendations api is deprecated, we looked to implement lb-radio support (https://git.asdf.cafe/abs3nt/gspot/pulls/54) which is working, but it is not great, since LB is not first-class.

this is a new project, which uses listenbrainz/musicbrainz as the primary source for music, and simply uses spotify permium as a source to play music.

this would allow more control over playback and metadata, along with doing things like scrobble directly to LB, play music from the LB radio, synchronize LB with spotify automatically, etc.

one big difference is that https://github.com/devgianlu/go-librespot now exists, which did not exist when we made gspot. so now, a gobrainsd + gobc client both make sense :)

## progress

currently, https://git.asdf.cafe/abs3nt/gspot/pulls/54 exists as an MVP

i am currently working on research aroudn the existing radio endpoints, with the hope to create a new endpoint that will allow for reccomendations based off a specific track.
