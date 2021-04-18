import React, { FunctionComponent } from 'react';
import SidebarContext from './SidebarContext';

type SidebarItemProps = {
  itemId: number;
  setActiveItem: (itemId: number) => void;
};

const SidebarItem: FunctionComponent<SidebarItemProps> = (
  props
): JSX.Element => {
  const toggle = (): void => {
    props.setActiveItem(props.itemId);
  };

  return (
    <SidebarContext.Consumer>
      {({ activeItem }) => {
        return (
          <div
            className={
              `flex items-center pl-12 space-x-6 cursor-pointer ` +
              `${
                activeItem === props.itemId
                  ? 'text-pink-600 border-pink-600 border-r-2 font-semibold'
                  : 'text-gray-600'
              }`
            }
            onClick={toggle}>
            {props.children}
          </div>
        );
      }}
    </SidebarContext.Consumer>
  );
};

export default SidebarItem;
