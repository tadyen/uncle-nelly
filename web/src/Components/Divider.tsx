import React from 'react';

export type DividerProps = {
    start?: React.ReactNode,
    middle?: React.ReactNode,
    end?: React.ReactNode,
    children?: React.ReactNode,
}
export default function Divider(prop: DividerProps){
    return (<>
<div className="flex w-full flex-col">
  <div className="divider divider-start">
    {prop.start}                
  </div>
  <div className="divider">
    {prop.middle}
  </div>
  <div className="divider divider-end">
    {prop.end}
  </div>
</div>
    </>)
}

