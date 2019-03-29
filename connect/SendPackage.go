package connect

import (
    "../models"
    "../settings"
)

func SendEncryptedPackage(pack settings.Package) int8 {
    if settings.User.NodeConnection[pack.To] != 1 {
        nullNode(pack.To)
        return settings.EXIT_FAILED
    }

    var new_pack = settings.Package {
        From: models.From {
            Name: pack.From.Name,
        },
        To: encrypt(settings.User.NodeSessionKey[pack.To], pack.To),
        Head: models.Head {
            Header: encrypt(settings.User.NodeSessionKey[pack.To], pack.Header),
            Mode: encrypt(settings.User.NodeSessionKey[pack.To], pack.Mode),
        },
        Body: encrypt(settings.User.NodeSessionKey[pack.To], pack.Body),
    }

    return sendNodePackage(pack.To, new_pack)
}