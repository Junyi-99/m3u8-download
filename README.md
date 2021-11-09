# m3u8-download

M3U8 downloader written in GO.

### Playlist

- Media Play List
    - a list of Media Segments

```text
#EXTM3U // format identifier tag
#EXT-X-TARGETDURATION:10 // 所有媒体分段都小于 10s

#EXTINF:9.009,  // 时长 9.009s
http://media.example.com/first.ts
#EXTINF:9.009,
http://media.example.com/second.ts
#EXTINF:3.003,
http://media.example.com/third.ts
```

- Master Play List
  - provides a set of Variant Streams, each of which
    describes a different version of the same content.
    - A Variant Stream 
      - a Media Playlist that specifies media encoded at a particular bit rate, in a particular format, and at a
        particular resolution for media containing video.
      - A Variant Stream can also specify a set of Renditions.  Renditions
        are alternate versions of the content, such as audio produced in
        different languages or video recorded from different camera angles.
        
      - Clients should switch between different Variant Streams to adapt to
        network conditions.  Clients should choose Renditions based on user
        preferences.


UTF-8 text files containing URIs and descriptive tags.

