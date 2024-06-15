import { style } from '@vanilla-extract/css'

export default {
  main: style({
    margin: '10px 0',
    position: 'relative',
    cursor: 'pointer',
    display: 'block',
    padding: '20px 10px',
    width: '90%',
    border: 'solid 1px rgba(255,255,255,0.3)',
    textAlign: 'left',
  }),
  created: style({
    position: 'absolute',
    bottom: '0',
    right: '0',
    fontSize: '13px',
    color: 'rgba(255,255,255,0.3)',
  }),
}
