
const niceSize = (size: number) => {
    let suffix = ""
    if (size < 1000000) {
        size = size / 1000
        suffix = "KB"
    } else if (size < 1000000000) {
        size = size / 1000000
        suffix = "MB"
    } else {
        size = size / 1000000000
        suffix = "GB"
    }
    return `${Math.round(size * 10)/10}${suffix}`
}

export default {
    niceSize
}