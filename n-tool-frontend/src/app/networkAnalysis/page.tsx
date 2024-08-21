//This page calls the scanner component UI which displays search results. Will also be the page where one could initiate a scan. Search results will eventually be displayed only after a scan is done.
import React from "react";
import ScanResults from "@/features/scanResults/scanResults";

const NetworkAnalysis: React.FC = () => {
  return (
    <div>
      <h1>Network Analysis</h1>
      <ScanResults />
    </div>
  );
};

export default NetworkAnalysis;
