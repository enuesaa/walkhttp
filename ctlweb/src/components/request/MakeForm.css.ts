import { style } from '@vanilla-extract/css'

export default {
  main: style({
    maxWidth: '700px',
    margin: '0 auto',
    fontSize: '25px',
    lineHeight: '1.7',
  }),
  from: style({
    textAlign: 'center',
  }),
  method: style({
    margin: '10px 0',
  }),
  url: style({
    margin: '10px 0',
  }),
  btn: style({
    margin: '30px 0',
    padding: '0 10px',
  })
}
