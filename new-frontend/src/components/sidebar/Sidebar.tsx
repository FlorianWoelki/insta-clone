import React, { FunctionComponent, ReactElement } from 'react';

type SidebarProps = {
  sidebarItems: ReactElement;
};

const Sidebar: FunctionComponent<SidebarProps> = (props): JSX.Element => {
  return (
    <div className="py-8 bg-gray-100 xl:col-span-2 2xl:col-span-1">
      {props.children}

      <div className="space-y-5">{props.sidebarItems}</div>
    </div>
  );
};

export default Sidebar;
