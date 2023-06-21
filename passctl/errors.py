class CommandNotFound(Exception):
    def __init__(self):
        super(CommandNotFound, self).__init__("Command not found")

class MasterInvalid(Exception):
    def __init__(self):
        super(MasterInvalid, self).__init__("Master password is invalid")

class ServerURLInvalid(Exception):
    def __init__(self):
        super(ServerURLInvalid, self).__init__("API server is not a valid passctl server")

class ServerPassInvalid(Exception):
    def __init__(self):
        super(ServerPassInvalid, self).__init__("API server password is not valid")

class ServerMaxData(Exception):
    def __init__(self):
        super(ServerMaxData, self).__init__("Server max vault limit exceeded")

class ServerNotSet(Exception):
    def __init__(self):
        super(ServerNotSet, self).__init__("Server is not set")

class ParameterError(Exception):
    def __init__(self):
        super(ParameterError, self).__init__("Invalid parameters")

class NotFoundInVault(Exception):
    def __init__(self):
        super(NotFoundInVault, self).__init__("Requested key is not found in the vault")

class AlreadyInVault(Exception):
    def __init__(self):
        super(AlreadyInVault, self).__init__("Requested key is already in the vault")
