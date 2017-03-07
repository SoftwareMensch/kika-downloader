# ABSTRACT
This is a simple tool to download offered videos from the official KIKA
website. KIKA is a trademark of the german tv stations ARD and ZDF.

# DISCLAIMER
Video download tool for the german, child-oriented, TV channel KIKA.
KIKA is a trademark of the german TV stations ARD and ZDF. You may not
be allowed to share downloaded content with the public without a
written permission from KIKA.

Please use this code only for educational purposes.

# EXAMPLE USAGE
Current commands require at least an entry URL to start processing. An example URL
is [this](http://www.kika.de/super-wings/sendungen/buendelgruppe2430_page-0_zc-f897245a_zs-de936bd8.html).

You can also provide a socks proxy, so it is possible to transfer traffic through the TOR network
or other SOCKS protocol based hosts.

## OPTIONS VERSION 1.0.0
```
Usage of kika-downloader:

Options:

  -socks-proxy-url=<socks5://127.0.0.1:9050>               optional socks proxy (i.e. TOR)

Commands:

  fetch-all -url=<entry url> -output-dir=<download dir>    download all videos to given directory
  print-csv -url=<entry url>                               print csv like output of all videos
```

## DOWNLOAD ALL EPISODES OF A SERIES
```
kika-downloader \
        fetch-all \
        -url=http://www.kika.de/mouk-der-weltreisebaer/sendungen/allevideosmoukderweltreisebaer100_page-0_zc-6615e895.html \
        -output-dir=<YOUR DOWNLOAD DIR>
```

## CSV OUTPUT OF ALL EPISODES OF A SERIES
```
kika-downloader \
        print-csv \
        -url=http://www.kika.de/mouk-der-weltreisebaer/sendungen/allevideosmoukderweltreisebaer100_page-0_zc-6615e895.html \
```

The optional parameter ```-socks-proxy-url``` is available for both commands.
