export function squarePreview() {
  let imgs = document.getElementsByClassName('preview')
  for (let img of imgs) img.style.height = `${img.clientWidth}px`
}

export function squarePreviewDiv() {
  let imgs = document.querySelectorAll('.preview')
  for (let img of imgs) {
    img.style.maxHeight = `${img.clientWidth}px`
    img.style.minHeight = `${img.clientWidth}px`
  }
}
