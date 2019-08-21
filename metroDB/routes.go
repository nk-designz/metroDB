package main

import (
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        getIndex,
    },
    Route{
        "Metrics",
        "GET",
        "/metrics",
        getMetrics,
    },
    Route{
        "HealthCheck",
        "GET",
        "/health",
        getHealthCheck,
    },
    Route{
        "doc Handler",
        "GET",
        "/doc",
        getApiTree,
    },
    Route{
        "API Handler",
        "GET",
        "/api",
        getStackList,
    },
    Route{
        "Stack Insight Handler",
        "GET",
        "/api/{stackName}/info",
        getStackInfo,
    },
    Route{
        "Get Block List Handler",
        "GET",
        "/api/{stackName}/all",
        getBlockList,
    },
    Route{
        "Block Get labels Handler",
        "GET",
        "/api/{stackName}/label",
        getLabels,
    },
    Route{
        "Block Set Label Handler",
        "PATCH",
        "/api/{stackName}/label",
        setLabel,
    },
    Route{
        "Stack Append Handler",
        "POST",
        "/api/{stackName}",
        addNewStack,
    },
    Route{
        "Stack New Block Handler",
        "PATCH",
        "/api/{stackName}",
        addNewBlock,
    },
    Route{
        "Stack Get Handler",
        "GET",
        "/api/{stackName}",
        getStack,
    },
    Route{
        "Block Get Handler",
        "GET",
        "/api/{stackName}/{blockId}",
        getBlock,
    },
    Route{
        "Block Get Raw Handler",
        "GET",
        "/api/{stackName}/{blockId}/raw",
        getBlockRaw,
    },
}
