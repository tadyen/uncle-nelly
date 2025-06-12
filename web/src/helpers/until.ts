export default function until(fn: () => boolean): Promise<void> {
    return new Promise( resolve => {
        const intervalCode = setInterval(()=>{
            if( fn()) {
                resolve()
                clearInterval(intervalCode)
            }
        })
    })
}
