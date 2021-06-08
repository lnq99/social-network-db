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

export function handlerResponse(response) {
  let responseOK =
    response && response.status === 200 && response.statusText === 'OK'
  if (responseOK) {
    let data = response.data
    console.log(data)
    if (data) {
      return data
    }
  }
}
