
  

# Project Directory Structure

  

## Project Root

  

-  `n-tool-frontend`

  

### `public`

  

####  `assets`

#### `testJson.json` - Sample JSON file for testing datatable `

#### `next.svg`

####  `vercel.svg`

  
---

### `src`

#### `app` - All folders inside with page.tsx are pages in Next.js

 
##### `history`

###### -`page.tsx`

#####  `home`

###### -`page.tsx`

#####  `login`

###### -`page.tsx`

#####  `scanner`

###### -`page.tsx`

##### `verifyEnv`
###### -  `page.tsx` - Test for environment variables


#####  `colors.css`

#####  `favicon.ico`

#####   `globals.css`

#####   `layout.tsx`

#####   `page.tsx`

  
#### `components`

#####  `footer`

#####  `header`

##### `scanResults`
###### -scanResults.tsx` - Component for displaying scan results

####   `side`

####  `ui`

###### -ui components` - SHA CDN & Custom UI components

#### `verifyEnv`

- `verifyEnv.tsx` - Test script to make sure env variables are being read

  

#### `lib`

#### `utils`

-  `debugLogUtil.ts` - Disables/Enables debug based on env
-  `folderRouter.ts` - Router utilities
-  `resultsParser.ts` - Parses the json data
-  `resultsUtils.ts` - Utilities for scanResults

 --- 

### Root Files

  

-  `components.json`

-  `next-env.d.ts`

-  `next.config.mjs`

-  `package-lock.json`

-  `package.json`

-  `postcss.config.mjs`

-  `README.md`

-  `tailwind.config.ts`

-  `tsconfig.json`

-  `.env.local`