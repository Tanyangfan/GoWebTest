#-*- coding: UTF-8 -*-

import sys
import os

def doBuild(file):
	if not os.path.exists('out'):
		os.mkdir('out')
	os.chdir('./out')
	os.system(('go build ../%s' % file))

def todoBuild():
	params = sys.argv
	if len(params) == 2:
		file = params[1]
		if '.go' in file:
			doBuild(file)
			#os.system(('go build %s' % file))
		else:
			print 'params error'
		
	elif len(params) == 3:
		file = params[1]
		os_type = params[2]
		if '.go' in file:
			out_file = file
			os.environ['CGO_ENABLED']='0'
			if os_type == 'win':
				os.environ['GOOS']='windows'
				os.environ['GOARCH']='amd64'
			elif os_type == 'linux':
				os.environ['GOOS']='linux'
				os.environ['GOARCH']='amd64'
			elif os_type == 'linux32':
				os.environ['GOOS']='linux'
				os.environ['GOARCH']='arm'
			doBuild(file)
		else:
			print 'params error'
	else:
		print 'params error'

if __name__ == "__main__":
	print sys.argv[1:]
	todoBuild()