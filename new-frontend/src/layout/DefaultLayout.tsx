import React, { FunctionComponent, useState } from 'react';
import Avatar from '../components/Avatar';
import Icon from '../components/Icon';
import Sidebar from '../components/sidebar/Sidebar';
import SidebarContext from '../components/sidebar/SidebarContext';
import SidebarItem from '../components/sidebar/SidebarItem';

type SidebarItemsProps = {
  setActiveItem: React.Dispatch<React.SetStateAction<number>>;
};

const SidebarItems: FunctionComponent<SidebarItemsProps> = ({
  setActiveItem,
}): JSX.Element => {
  return (
    <React.Fragment>
      <SidebarItem itemId={0} setActiveItem={() => setActiveItem(0)}>
        <Icon name="template" className="w-5 h-5" />
        <p>Feed</p>
      </SidebarItem>
      <SidebarItem itemId={1} setActiveItem={() => setActiveItem(1)}>
        <Icon name="search" className="w-5 h-5" />
        <p>Explore</p>
      </SidebarItem>
      <SidebarItem itemId={2} setActiveItem={() => setActiveItem(2)}>
        <Icon name="bell" className="w-5 h-5" />
        <p>Notifications</p>
      </SidebarItem>
      <SidebarItem itemId={3} setActiveItem={() => setActiveItem(3)}>
        <Icon name="mail" className="w-5 h-5" />
        <p>Messages</p>
        <span className="text-sm text-gray-400">8</span>
      </SidebarItem>
      <SidebarItem itemId={4} setActiveItem={() => setActiveItem(4)}>
        <Icon name="paper-airplane" className="w-5 h-5" />
        <p>Direct</p>
      </SidebarItem>
      <SidebarItem itemId={5} setActiveItem={() => setActiveItem(5)}>
        <Icon name="chart-bar" className="w-5 h-5" />
        <p>Stats</p>
      </SidebarItem>
      <SidebarItem itemId={6} setActiveItem={() => setActiveItem(6)}>
        <Icon name="cog" className="w-5 h-5" />
        <p>Settings</p>
      </SidebarItem>

      <div className="flex items-center">
        <div className="flex-grow bg-gray-200 border"></div>
        <div className="p-1 -mr-2 bg-white rounded-full">
          <Icon name="chevron-left" className="w-5 h-5" />
        </div>
      </div>

      <div className="flex items-center pl-12 space-x-6 text-gray-600">
        <Icon name="logout" className="w-5 h-5" />
        <p>Logout</p>
      </div>
    </React.Fragment>
  );
};

type DefaultLayoutProps = {
  defaultActiveItem: number;
};

const DefaultLayout: FunctionComponent<DefaultLayoutProps> = (
  props,
): JSX.Element => {
  const [activeItem, setActiveItem] = useState(props.defaultActiveItem);

  return (
    <div className="grid h-screen grid-cols-4 xl:grid-cols-7 2xl:grid-cols-5">
      <SidebarContext.Provider value={{ activeItem }}>
        <Sidebar sidebarItems={<SidebarItems setActiveItem={setActiveItem} />}>
          <div className="px-8 space-y-3">
            <Avatar />
            <div className="space-y-0">
              <h2 className="text-lg font-bold text-center text-gray-800">
                Kate Lingard
              </h2>
              <p className="text-sm text-center text-gray-400">@katy69</p>
            </div>
          </div>

          <div className="grid justify-between grid-cols-3 gap-6 px-8 my-8 text-sm">
            <div className="text-center">
              <p className="font-bold text-gray-800">46</p>
              <p className="text-gray-400">Posts</p>
            </div>
            <div className="text-center">
              <p className="font-bold text-gray-800">2.8k</p>
              <p className="text-gray-400">Followers</p>
            </div>
            <div className="text-center">
              <p className="font-bold text-gray-800">526</p>
              <p className="text-gray-400">Following</p>
            </div>
          </div>
        </Sidebar>
      </SidebarContext.Provider>

      <div className="container col-span-3 px-16 mx-auto my-12 xl:col-span-5 2xl:col-span-4">
        {props.children}
      </div>
    </div>
  );
};

export default DefaultLayout;
