(async () => {
    interface Node {
        text: string;
        children: Node[];
    }

    let iRootNode: Node = {} as Node
    let stSeparator: string = '\\'
    const gstPaths: string[] = []

    const loadJson = () => {
        return new Promise((resolve, reject) => {
            const xhr = new XMLHttpRequest()
            xhr.open('GET', 'dir.json', true)
            xhr.onreadystatechange = () => {
                if (4 != xhr.readyState) {
                 return
                }

                if (200 != xhr.status) {
                    // reject()
                    throw 'Load JSON Error'
                }

                iRootNode = JSON.parse(xhr.responseText)
                resolve('')
            }
            xhr.send(null)
        })
    }

    const parseNode = (iNode: Node, stParentDir: string) => {
        if (iNode.text) {
            createDir(iNode, stParentDir)
        }

        if (stParentDir) {
            stParentDir += stSeparator
        }

        iNode.text && (stParentDir += iNode.text)

        if (!iNode.children) {
            return
        }

        for (const iChildNode of iNode.children) {
            parseNode(iChildNode, stParentDir)
        }
    }

    const createDir = (iNode: Node, stParentDir: string) => {
        let path = iNode.text
        if (stParentDir) {
            path = stParentDir + stSeparator + path
        }

        gstPaths.push(path)
    }

    const generateBatFile = () => {
        let stResult = ''
        for (const path of gstPaths) {
            stResult += `md ${path} \r\n`
        }
        const url: string = URL.createObjectURL(new Blob([stResult], {type: 'text/plain'}))

        const domA: HTMLElement = document.createElement('a')
        domA.setAttribute('href', url)
        domA.setAttribute('target', '_blank')
        domA.setAttribute('download', 'generate_dir_03.bat')
        document.body.appendChild(domA)
        domA.click()
        setTimeout(() => {
            document.body.removeChild(domA)
        }, 500)
    }

    await loadJson()
    parseNode(iRootNode, '')
    generateBatFile()
})()