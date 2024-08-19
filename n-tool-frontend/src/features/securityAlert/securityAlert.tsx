"use client";

import * as React from "react";
import { Alert, AlertTitle, AlertDescription } from "@/components/ui/alert";
import { Tooltip, TooltipTrigger, TooltipContent, TooltipProvider } from "@/components/ui/tooltip";
import { cn } from "@/lib/utils";

interface SecurityAlertProps {
  alertCount: number;
}

const SecurityAlert: React.FC<SecurityAlertProps> = ({ alertCount }) => {
  // Determine the icon color based on the alertCount
  const getIconColor = () => {
    if (alertCount > 80) return "text-red-500";
    if (alertCount > 40) return "text-yellow-500";
    return "text-green-500";
  };

  return (
    <TooltipProvider>
      <Alert className="relative w-full rounded-lg border-none  p-4 ${isRootMode ? 'bg-gray-800 text-slate-50' : 'bg-slate-50 text-gray-800'}">
        <Tooltip>
          <TooltipTrigger asChild>
            <i className={cn("fa-solid fa-triangle-exclamation ", getIconColor(), "absolute left-4 top-4 text-xl")} />
          </TooltipTrigger>
          <TooltipContent side="top">
            Security risks are rated by color:
            <br />
            <strong>Red</strong>: High risk level
            <br />
            <strong>Yellow</strong>: Medium risk level
            <br />
            <strong>Green</strong>: Low risk level
          </TooltipContent>
        </Tooltip>

        <div className="pl-12 ">
          <AlertTitle>Security Alerts</AlertTitle>
          <AlertDescription>
            You have {alertCount} security risks.
          </AlertDescription>
        </div>
      </Alert>
    </TooltipProvider>
  );
};

export default SecurityAlert;
