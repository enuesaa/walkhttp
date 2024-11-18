export const judgeStatusColor = (status: number): 'green' | 'red' | 'blue' => {
  if (200 <= status && status < 400) {
    return 'green'
  }
  if (400 <= status && status < 600) {
    return 'red'
  }
  return 'blue'
}
