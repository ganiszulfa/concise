# Concise Web

## Tech Specs

1. Svelte
1. NPM
1. Tailwind
1. [Daisy UI](https://github.com/saadeghi/daisyui)
1. JSDoc (tba)

## Local Deployment

1. Run the concise backend if you want to use your local backend if not, adjust `src/lib/configs.js` accordingly.
2. install necessary packages

```
npm install
```

3. Run the Svelte Web

```
npm run dev --host=0.0.0.0
```

4. (optional) If you are editing the design/css, run this in another terminal:

```
npm run csswatch
```

## Cloud Deployment

1. Concise Web has been tailored for deployment on [CapRover](https://caprover.com/). CapRover is an amazing piece of software, it's like Heroku of your own! Install this on your own server.
    - Add a custom app in CapRover (the name will be WEB_APP_NAME).
    - Set deployment method to "Official CLI", generate app token (it will be WEB_APP_TOKEN).

2. Go to your `Secrets and variables` section in your Github Repo. Add these variables
    - WEB_CAPROVER_SERVER
    - WEB_APP_NAME
    - WEB_APP_TOKEN

3. Push a tag with this format `wv[0-9]+.[0-9]+.[0-9]+`, for example: `wv1.0.2`