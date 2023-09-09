import { css } from '@emotion/react'
import normalize from 'emotion-normalize'

export const globalStyle = css`
  ${normalize}

  html,
  body {
    padding: 0;
    margin: 0;
    font-family: -apple-system, BlinkMacSystemFont, Segoe UI, sans-serif;
  }
  a {
    color: inherit;
    text-decoration: none;
  }
  * {
    box-sizing: border-box;
    letter-spacing: 0.9px;
    word-wrap: break-word;
  }
  body {
    background: #1d2029;
  }

  input, textarea {
    background: inherit;
    color: inherit;
    outline: none;
    appearance: none;
    border: none;
    display: block;
  }

  button {
    outline: none;
    appearance: none;
    border: none;
    display: block;
    cursor: pointer;
  }

  ul {
    list-style-type: none;
    padding-inline-start: 0;
  }

  h1 {
    margin: 0;
  }
`