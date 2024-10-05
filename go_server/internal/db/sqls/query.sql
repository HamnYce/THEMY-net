-- name: GetHost :one
SELECT
    *
FROM
    hosts
WHERE
    id = ?
LIMIT
    1;

-- name: ListHosts :many
SELECT
    *
FROM
    hosts
LIMIT
    ?
OFFSET
    ?;

-- name: CreateHost :one
INSERT INTO
    hosts (
        name,
        mac,
        ip,
        hostname,
        status,
        exposure,
        internetAccess,
        os,
        osVersion,
        ports,
        usage,
        location,
        owners,
        dependencies,
        createdAt,
        createdBy,
        recordedAt,
        access,
        connectsTo,
        hostType,
        exposedServices,
        cpuCores,
        ramGB,
        storageGB
    )
VALUES
    (
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?
    ) RETURNING id;

-- name: UpdateHost :exec
UPDATE hosts
SET
    name = ?,
    mac = ?,
    ip = ?,
    hostname = ?,
    status = ?,
    exposure = ?,
    internetAccess = ?,
    os = ?,
    osVersion = ?,
    ports = ?,
    usage = ?,
    location = ?,
    owners = ?,
    dependencies = ?,
    createdAt = ?,
    createdBy = ?,
    recordedAt = ?,
    access = ?,
    connectsTo = ?,
    hostType = ?,
    exposedServices = ?,
    cpuCores = ?,
    ramGB = ?,
    storageGB = ?
WHERE
    id = ?;

-- name: DeleteHost :one
DELETE FROM hosts
WHERE
    id = ? RETURNING *;