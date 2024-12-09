pagination : start to end ( find out till how many pages preceeding is the data present)
backfilling : going back to get data to fill data
polling : will fetch data from today to future
APIFY : tol for scraping

in round 1:
    for each page , accumulate all the links available(just the links) and create a map

in round 2:
    for each link , scrape the content 

potential edge case:
    artice get deleted 
        implement a **wait timer** ( to delay if fetching unsuccessful)
        if after K times , article not fetched -> put into **backlog table** 
        also prepare a Potentially deleetd or **archive table**
    
    new article added     
        polling added -> prepare a data for timer , lastpage , **lastarticle**

    modified
    article contains contents extra than wanted 
    chnage in html structure
        APIFY


potential solutions :
    

in round 3:
    special purpose for scraped data like sentiment analysis

in round 4:
    human intervention and approval 


** also add a analytics and feedback section (conditional)