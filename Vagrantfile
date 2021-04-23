$script = <<-SCRIPT

# echo "Preparing local node_modules folder…"
# mkdir -p /home/vagrant/app/sdk/vagrant_node_modules
# chown vagrant:vagrant /home/vagrant/app/sdk/vagrant_node_modules

echo "cd /vagrant/go/src/app" >> /home/vagrant/.profile
echo "cd /vagrant/go/src/app" >> /home/vagrant/.bashrc
echo "All good!!"
SCRIPT

VAGRANTFILE_API_VERSION = "2"

# Vagrant base box to use
BOX_BASE = "bento/ubuntu-20.04"
# amount of RAM for Vagrant box
BOX_RAM_MB = "3072"
# number of CPUs for Vagrant box
BOX_CPU_COUNT = "1"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
    
    # 1. Use this for "Standard setup"
    config.vm.box = BOX_BASE
    # Uncomment the lines below if you would like to protect the VM
    # config.ssh.username = 'vagrant'
    # config.ssh.password = 'vagrant'
    # config.ssh.insert_key = 'true'

    config.vm.synced_folder ".", "/vagrant/go/src/app"

    # Ports forward
    # For Api
    config.vm.network "forwarded_port", guest: 9000, host: 9000

    #Mongodb
    config.vm.network "forwarded_port", guest: 27017, host: 27017
    

    # This gets executed for both vm1 & vm2
    # config.vm.provision "shell", inline:  "echo 'All good'"
    config.vm.provision "shell", inline:  $script

    # config.vm.provision "shell", run: "always", inline: <<-SHELL
    #   mount --bind  /home/vagrant/app/sdk/node_modules /home/vagrant/app/sdk/vagrant_node_modules
    # SHELL
  
    # To use a diffrent Hypervisor create a section config.vm.provider
    # And comment out the following section
    # Configuration for Virtual Box
    config.vm.provider :virtualbox do |vb|
      # Change the memory here if needed - 3 Gb memory on Virtual Box VM
      vb.customize ["modifyvm", :id, "--memory", BOX_RAM_MB, "--cpus", BOX_CPU_COUNT]
      # Change this only if you need destop for Ubuntu - you will need more memory
      vb.gui = false
      # In case you have DNS issues
      # vb.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
    end

    # Configuration for Windows Hyperv
    config.vm.provider :hyperv do |hv|
      # Change the memory here if needed - 2 Gb memory on Virtual Box VM
      hv.customize ["modifyvm", :id, "--memory", BOX_RAM_MB, "--cpus", BOX_CPU_COUNT]
      # Change this only if you need destop for Ubuntu - you will need more memory
      hv.gui = false
      # In case you have DNS issues
      # vb.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
    end


  end