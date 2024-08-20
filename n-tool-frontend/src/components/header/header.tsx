"use client";

import React, { useState, useEffect } from "react";
import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
} from "@/components/ui/navigation-menu";
import { Input } from "@/components/ui/input";
import { Switch } from "@/components/ui/switch";
import SecurityAlert from "@/features/securityAlert/securityAlert";
import "./header.css";

interface HeaderProps {
  orderedFolders: { name: string; index: number }[];
}

export default function Header({ orderedFolders }: HeaderProps) {
  const [isDarkMode, setIsDarkMode] = useState(false);

  useEffect(() => {
    const rootElement = document.documentElement;
    if (isDarkMode) {
      rootElement.classList.add("dark");
    } else {
      rootElement.classList.remove("dark");
    }
  }, [isDarkMode]);

  const toggleTheme = () => {
    setIsDarkMode((prevMode) => !prevMode);
  };

  return (
    <header className="header">
      <nav className="nav-container">
        <NavigationMenu>
          <NavigationMenuList className="nav-list w-full">
            {orderedFolders.map((folder) => (
              <NavigationMenuItem key={folder.name} className="nav-item">
                <NavigationMenuLink
                  href={`/${folder.name}`}
                  className="nav-link"
                >
                  {folder.name.charAt(0).toUpperCase() + folder.name.slice(1)}
                </NavigationMenuLink>
              </NavigationMenuItem>
            ))}
          </NavigationMenuList>
        </NavigationMenu>

        <div className="flex items-center space-x-4">
          <SecurityAlert alertCount={85} />
          <div className="dark-light-switch">
            <Switch checked={isDarkMode} onCheckedChange={toggleTheme} />
          </div>
        </div>
      </nav>
    </header>
  );
}
