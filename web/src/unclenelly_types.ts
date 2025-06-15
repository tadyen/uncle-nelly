export interface UncleNellyResult<T> {
    error?: string
    response?: T
}

export interface UncleNelly {
    init_job: () => UncleNellyResult<string|null>
    reset_product: () => UncleNellyResult<string|null>
    get_tables: () => UncleNellyResult<object|null>
    set_product_base: (productBase: string) => UncleNellyResult<string|null>
    cook_with: (...ingredients: string[]) => UncleNellyResult<string|null>
    product_info: () => UncleNellyResult<object|null>
}

// Expected type for get_tables response from the wasm
export type UncleNellyGetTables = {
    base_ingredients: {
        [key: string]: {
            Effect: string[],
            Name: string,
            Price: number,
            Type: string,
        }
    },
    effects: {
        [key: string]: {
            Name: string,
            Multiplier: number,
            Conversion: {
                [key: string]: string
            }[]
        }
    },
    mix_ingredients: {
        [key: string]: {
            Name: string,
            Effect: string,
            Price: number,
        }
    }
}

export type UncleNellyTables = UncleNellyGetTables & {
    base_ingredients: {
        [key: string]: {
            Icon?: string,
        }
    },
    mix_ingredients: {
        [key: string]: {
            Icon?: string,
        }
    },
    effects: {
        [key: string]: {
            Color?: string,
        }
    }
}
