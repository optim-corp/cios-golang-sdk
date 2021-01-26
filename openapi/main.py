# coding: utf-8
import yaml, os, codecs

all_openapi = './openapi.yml'
with codecs.open(all_openapi, 'r', encoding="utf-8") as file:
    yml = yaml.load(file,  Loader=yaml.FullLoader)
    yml["paths"] = {}
    for path in os.listdir('./openapis'):
        print('openapis/'+path)
        with codecs.open('./openapis/'+path, encoding="utf-8") as openapi:
            _yml = yaml.load(openapi, Loader=yaml.FullLoader)
            yml['paths'].update(_yml['paths'])

    with codecs.open(all_openapi, 'w', encoding="utf-8") as file:
        yaml.dump(yml, file, allow_unicode=True)

with codecs.open(all_openapi, 'r', encoding="utf-8") as file:
    body = file.read()
    body = body.replace('../components', './components')
    with codecs.open(all_openapi, 'w', encoding="utf-8") as file:
        file.write(body)
