const puppeteer = require('puppeteer')

// Or import puppeteer from 'puppeteer-core';

// Launch the browser and open a new blank page

const main = async () => {
    console.log("Opening page!")
    const browser = await puppeteer.launch({headless: true, args: ['--no-sandbox', 'disable-setuid-sandbox']});
    const page = await browser.newPage();
    
    // Navigate the page to a URL.
    await page.goto('http://k8s-default-frontend-b914e7a1c1-905065450.us-west-1.elb.amazonaws.com/');
    
    await page.setViewport({width: 1080, height: 1024});
    
    setTimeout(async function() {
        await browser.close();
        main()
    }, 15000)
}

main()