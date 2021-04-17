import React, { FunctionComponent } from 'react';

type SidebarProps = {};

const Sidebar: FunctionComponent<SidebarProps> = (props): JSX.Element => {
  return <div>{props.children}</div>;
};

export default Sidebar;
