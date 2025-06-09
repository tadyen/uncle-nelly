// This file loads WASM implementaion of UncleNelly

import '../../wasm_exec';
import wasm from '../../main.wasm?url';
import { UncleNelly } from './unclenelly_types';


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

    const go_run = async () => {
        const go = new window.Go();
        const result = await WebAssembly.instantiateStreaming(
            fetch(wasm),
            go.importObject
        )
        await go.run(result.instance);
        return // instance crashed or terminated
    }

    const go_runner = async () => {
        while(true){
            await go_run();
            console.log('go instance crashed, restarting ...');
        }
    }

    go_runner();

    // wait until WASM create the function
    await until(() => window.InitUncleNelly != undefined)
    console.log('Uncle Nelly WASM loaded')
    return window.InitUncleNelly
}
