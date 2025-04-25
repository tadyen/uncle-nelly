// This file loads WASM implementaion of UncleNelly

import '../../wasm_exec'
import wasm from '../../main.wasm?url'

type stringOrNull = string | null

interface UncleNellyResult<T> {
    result?: T
    error?: string
}

declare global {
    export interface Window {
        Go: {
            new (): {
                run: (inst: WebAssembly.Instance) => Promise<void>,
                importObject: WebAssembly.Imports,
            }
        }
        InitUncleNelly: () => UncleNelly,
    }
}

export interface UncleNelly {
    init_job: () => UncleNellyResult<stringOrNull>
    reset_product: () => UncleNellyResult<stringOrNull>
    get_tables: () => UncleNellyResult<object|null>
    set_product_base: (productBase: string) => UncleNellyResult<stringOrNull>
    cook_with: (ingredients: string[]) => UncleNellyResult<stringOrNull>
    product_info: () => UncleNellyResult<object|null>
}

const until = (f: () => boolean): Promise<void> => {
    return new Promise(resolve => {
        const intervalCode = setInterval(() => {
            if (f()) {
                resolve()
                clearInterval(intervalCode)
            }
        }, 10)
    })
}

export async function loadUncleNelly() {
    if (!WebAssembly.instantiateStreaming) {
        // polyfill
        WebAssembly.instantiateStreaming = async (resp, importObject) => {
            const source = await (await resp).arrayBuffer()
            return await WebAssembly.instantiate(source, importObject)
        }
    }

    console.log('Loading WebAssembly')

    if (!WebAssembly) {
        throw new Error('WebAssembly is not supported in your browser')
    }
    const go = new window.Go()
    const result = await WebAssembly.instantiateStreaming(
        fetch(wasm),
        go.importObject
    )

    console.log('Loaded WebAssembly')

    go.run(result.instance)

    // wait until WASM create the function
    await until(() => window.InitUncleNelly != undefined)
    return window.InitUncleNelly
}
