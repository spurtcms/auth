-- Must use separate insert queries for parent and child data insertions to maintain sequence.

INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Default Category', 'default_category', 'time', uid, 0, 0, 'Default_Category', tid);

INSERT INTO tbl_categories(category_name, category_slug, created_on, created_by,is_deleted, parent_id, description, tenant_id)	VALUES ('Default1', 'default1', 'time', uid, 0, pid, 'Default1',tid);

INSERT INTO tbl_channels(channel_name, slug_name, field_group_id, is_active, is_deleted, created_on, created_by,channel_description,channel_type,tenant_id) VALUES ('Default_Channel', 'default_channel', 0, 1, 0, 'time', uid, 'default description','mychannels',tid);

INSERT INTO tbl_channel_categories(channel_id, category_id, created_at, created_on,tenant_id) VALUES (1, 'mapcat', 1, 'time',tid);

INSERT INTO tbl_member_groups(name,slug,description,is_active,is_deleted,created_on,created_by,tenant_id) VALUES ('Default Group', 'default-group', '', 1,0, 'time', uid,tid);

INSERT INTO tbl_blocks(title, block_content, cover_image, is_active , prime , created_by, created_on,tenant_id) VALUES ('layout1','<div class="grid grid-cols-4 gap-[8px]"><label class="bg-[#F4F4F4] flex items-center justify-center cursor-pointer aspect-[1/1] row-span-2 col-span-2 gap-[8px]"><img src="/public/img/upload-folder.svg" class="max-w-full max-h-full object-contain" alt="Upload-image">Upload Image</label><label class="bg-[#F4F4F4] flex items-center justify-center cursor-pointer aspect-[1/1] gap-[8px]"><img src="/public/img/upload-folder.svg" class="max-w-full max-h-full object-contain" alt="Upload-image">Upload Image</label><label class="bg-[#F4F4F4] flex items-center justify-center cursor-pointer aspect-[1/1] gap-[8px]"><img src="/public/img/upload-folder.svg" class="max-w-full max-h-full object-contain" alt="Upload-image">Upload Image </label><label class="bg-[#F4F4F4] flex items-center justify-center cursor-pointer aspect-[1/1] gap-[8px]"><img src="/public/img/upload-folder.svg" class="max-w-full max-h-full object-contain" alt="Upload-image">Upload Image</label><label class="bg-[#F4F4F4] flex items-center justify-center cursor-pointer aspect-[1/1] gap-[8px]"><img src="/public/img/upload-folder.svg" class="max-w-full max-h-full object-contain" alt="Upload-image">Upload Image </label></div>','/image-resize?name=media%2flayout1.png',1,0,uid,'time',tid),('layout2','<div class="grid grid-cols-2 gap-[8px]"><label class="bg-[#F4F4F4] flex items-center justify-center cursor-pointer h-[171px] w-full gap-[8px]"><img src="/public/img/upload-folder.svg" class="max-w-full max-h-full object-contain" alt="Upload-image">Upload Image</label><label class="bg-[#F4F4F4] flex items-center justify-center cursor-pointer h-[171px] w-full gap-[8px]"><img src="/public/img/upload-folder.svg" class="max-w-full max-h-full object-contain" alt="Upload-image">Upload Image </label><div class="col-span-2"><p  class="text-[16px] font-[400] leading-[22px] tracking-[0.005em] text-left text-[#555555]">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p></div><div class="col-span-2"><p  class="text-[16px] font-[400] leading-[22px] tracking-[0.005em] text-left text-[#555555]"> Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p></div></div>','/image-resize?name=media%2flayout2.png',1,0,uid,'time',tid),('layout3','<div class="grid grid-cols-2 gap-[8px]"><label class="bg-[#F4F4F4] flex items-center justify-center cursor-pointer h-[171px] w-full col-span-2 gap-[3px]"><img src="/public/img/upload-folder.svg" class="max-w-full max-h-full object-contain" alt="Upload-image">Upload Image </label><label class="bg-[#F4F4F4] flex items-center justify-center cursor-pointer h-[171px] w-full gap-[8px]"><img src="/public/img/upload-folder.svg" class="max-w-full max-h-full object-contain" alt="Upload-image">Upload Image </label><label class="bg-[#F4F4F4] flex items-center justify-center cursor-pointer h-[171px] w-full gap-[8px]"><img src="/public/img/upload-folder.svg" class="max-w-full max-h-full object-contain" alt="Upload-image">Upload Image </label></div>','/image-resize?name=media%2flayout3.png',1,0,uid,'time',tid);

INSERT INTO tbl_blocks(title, block_content, cover_image, is_active , prime , created_by, created_on,tenant_id) VALUES('layout4','<div class="card-1 bg-white rounded flex items-center gap-[6px]"><img class="w-[140px] h-[156px] object-contain" src="/public/img/block-1.png"><div class="flex flex-col items-start gap-[8px] pe-[6px]"><h5 class="text-[#262626] font-semibold text-left text-sm mb-0">Lorem ipsum it i dolor sit amet</h5><p class="text-[#717171] font-normal mb-0 text-xs text-left">“Lorem ipsum dolor sit amet, consect it adipiscing elit, sed do eiusmod tempor Lorem ipsum dolor sitamet.”</p></div></div>','/blocks/IMG-1726547614.jpeg',1,0,uid,'time',tid),('layout5','<div class="card-2 bg-white rounded grid grid-cols-2 gap-[4px] p-[8px]"><div class="flex flex-col gap-[6px] items-start"><img class="max-h-[535px] max-w-[359px] object-contain" src="/public/img/block-2.png"><div class="flex flex-col gap-[4px] items-start"><h3 class="text-[#262626] text-[16px] font-normal">Lorem ipsum</h3><p class="text-[#717171] text-[14px] font-normal mb-0">“Lorem ipsum dolor sit amet, consectetur, seddo eiusmod tempor. Lorem ipsum dolor sit amet”</p></div></div><div class="flex flex-col gap-[6px] items-start"><img class="max-h-[535px] max-w-[359px] src="/public/img/block-2.png"><div class="flex flex-col gap-[4px] items-start"><h3 class="text-[#262626] text-[16px] font-normal mb-0">Lorem ipsum</h3><p class="text-[#717171] text-[14px] font-normal mb-0">“Lorem ipsum dolor sit amet, consectetur, seddo eiusmod tempor. Lorem ipsum dolor sit amet”</p></div></div></div>','/blocks/IMG-1726548909.jpeg',1,0,uid,'time',tid),('layout6','<div class="card-3 flex flex-col items-start rounded-[6px] bg-white "><img class="max-h-[1090px] max-w-[388px] object-contain" src="/public/img/block-3.png" class="w-full object-cover"><div class="p-[8px] flex flex-col gap-[6px] items-start"><h3 class="text-[#262626] text-[16px] font-normal mb-0">Lorem ipsum dolor sit amet, consectetur Lorem ipsum dolor sit amet, consectetur.</h3><p class="text-[#555555] text-[14px] font-normal mb-0">Lorem ipsum dolor sit amet, consectetur, seddo eiusmod tempor.</p></div></div>','/blocks/IMG-1726549463.jpeg',1,0,uid,'time',tid)

INSERT INTO tbl_blocks(title, block_content, cover_image, is_active , prime , created_by, created_on,tenant_id) VALUES('layout7','<div class="card-4 p-4 flex flex-col gap-[8px] bg-white rounded"><h4 class="text-sm font-normal text-[#717171] mb-0 text-center">Lorem ipsum</h4><h6 class="text-sm text-[#262626] font-normal mb-0 text-center">Lorem ipsum dolor Lorem</h6><p class="text-[#717171] font-normal text-xs text-center">“Lorem ipsum dolor sit amet, consectetur, sed do eiusmod dolor sit amet Lorem ipsum.”</p><a href="#" class="no-underline text-[16px] font-normal text-[#0F62FE] text-center">Learn more</a></div>','/blocks/IMG-1726549594.jpeg',1,0,uid,'time',tid),('layout8','<div class="card-5 p-[12px] bg-white rounded flex gap-[8px]"><div class="flex flex-col items-start gap-[6px]"><div class="w-[80px]"><img src="/public/img/block-4.png" alt="" class="w-full rounded-[2px]"></div><div class="flex gap-[4px]"><div class="w-[24px]"><img src="/public/img/block-4.png" alt=""></div><div class="w-[24px]"><img src="/public/img/block-4.png" alt=""></div></div></div><div class="flex flex-col gap-[8px] items-start w-full"><div class="flex flex-col items-start gap-[4px]"><h5 class="text-xs font-normal mb-0 text-[#262626]">Lorem ipsum dolor sit</h5><div class="flex gap-[1px] items-center"><img src="/public/img/reviewstar-full.svg" alt=""><img src="/public/img/reviewstar-full.svg" alt=""><img src="/public/img/reviewstar-full.svg" alt=""><img src="/public/img/reviewstar-full.svg" alt=""><img src="/public/img/reviewstar.svg" alt=""></div></div><p class="m-0 text-[#262626] text-xs font-normal">$ 15.00</p><div class="flex flex-col items-start gap-[4px]"><h5 class="text-[14px] font-normal mb-0 text-[#717171]">Color</h5><div class="flex items-center gap-[1px]"><div class="w-[20px] h-[20px] rounded-full p-[2px] border border-[#000000]"><div class="bg-[#000000] h-full rounded-full"></div></div><div class="w-[20px] h-[20px] rounded-full p-[2px]"><div class="bg-[#004DFF] h-full rounded-full"></div></div><div class="w-[20px] h-[20px] rounded-full p-[2px]"><div class="bg-[#3B8620] h-full rounded-full"></div></div></div></div><div class="flex flex-col items-start gap-[4px]"><h5 class="text-[14px] font-normal mb-0 text-[#717171]">Size</h5><div class="flex items-center gap-[8px]"><div class="h-[16px] flex items-center justify-center px-[6px] border border-[#EDEDED] text-[14px] text-[#262626] font-normal rounded-[4px] uppercase">s</div><div class="h-[16px] flex items-center justify-center px-[6px] border bg-[#EBEBEB] border-[#EDEDED] text-[14px] text-[#262626] font-normal rounded-[4px] uppercase">m</div><div class="h-[16px] flex items-center justify-center px-[6px] border border-[#EDEDED] text-[14px] text-[#262626] font-normal rounded-[4px] uppercase">l</div><div class="h-[16px] flex items-center justify-center px-[6px] border border-[#EDEDED] text-[14px] text-[#262626] font-normal rounded-[4px] uppercase">xl</div></div></div></div></div>','/blocks/IMG-1726549783.jpeg',1,0,uid,'time',tid),('layout9','<div class="card-7 p-[8px] bg-white rounded grid grid-cols-2 gap-[20px]"><div class="flex flex-col gap-[8px] items-start"><div class="flex items-center gap-[8px]"><img class="w-[24px] h-[24px] object-contain" src="/public/img/block-5.svg"><h4 class="text-[#2F2F2F] font-medium text-xs mb-0">Lorem ipsum</h4></div><p class="text-[#555555] text-[16px] font-normal m-0">Lorem ipsum dolor sit amet, consectetur, sed do eiusmod tempor Lorem ipsum dolor sit amet, consectetur, sed do eiusmod tempor.</p></div><div class="flex flex-col gap-[8px] items-start"><div class="flex items-center gap-[8px]"><img class="w-[24px] h-[24px] object-contain" src="/public/img/block-5.svg"><h4 class="text-[#2F2F2F] font-medium text-xs mb-0">Lorem ipsum</h4></div><p class="text-[#555555] text-[16px] font-normal m-0">Lorem ipsum dolor sit amet, consectetur, sed do eiusmod tempor Lorem ipsum dolor sit amet, consectetur, sed do eiusmod tempor.</p></div></div>','/blocks/IMG-1726550244.jpeg',1,0,uid,'time',tid)

INSERT INTO tbl_blocks(title, block_content, cover_image, is_active , prime , created_by, created_on,tenant_id) VALUES('layout10','<div class="card-10 bg-white rounded p-[12px] flex gap-[8px]"><div class="w-[40%] bg-[#F9F9F9] p-[6px] flex flex-col gap-[4px] items-start"><p class="text-[#555555] text-[14px] font-normal m-0 text-start">Lorem ipsum dolor sit amet, consectetur,    sed do eiusmod tempor.Lorem ipsum dolor sit amet, consectetur. Lorem ipsum dolor sit amet.</p><div class="flex flex-col items-start"><h4 class="text-[14px] text-[#262626] font-medium mb-0">Product Manager</h4><p class="text-[18px] text-[#555555] font-normal m-0">Lorem ipsum</p></div></div><div class="w-[60%] p-[6px] flex shadow-card flex-col gap-[4px] items-start"><img class="max-h-[328px] max-w-[180px] object-contain" src="/public/img/block-9.png" alt="" class="w-full"><p class="text-[16px] font-normal text-[#2F2F2F] m-0">Lorem ipsum dolor sit amet, consectetur Lorem ipsum dolor sit amet, consectetur.</p></div></div>','/blocks/IMG-1726551883.jpeg',1,0,uid,'time',tid),('layout11','<div class="card-12 p-[8px] bg-white rounded grid grid-cols-2 gap-[8px]"><div class="flex flex-col gap-[6px] items-start"><div class="flex items-center flex-col gap-[4px]"><img class="w-[24px] h-[24px] object-contain" src="/public/img/block-5.svg"><h4 class="text-[#2F2F2F] font-medium text-xs mb-0">Lorem ipsum</h4></div><p class="text-[#555555] text-[16px] font-normal m-0 pb-[8px] border-b border-[#ECECEC] w-full">Lorem ipsum dolor sit amet, consectetur, sed do eiusmod tempor.</p><div class="grid grid-cols-2 gap-[16px] w-full"><div class="flex flex-col items-center"><h3 class="text-center text-[#2F2F2F] text-[16px] mb-0">75%</h3><p class="text-center text-[#555555] text-[18px] mb-0">Reduced monthly</p></div><div class="flex flex-col items-center"><h3 class="text-center text-[#2F2F2F] text-[16px] mb-0">75%</h3><p class="text-center text-[#555555] text-[18px] mb-0">Reduced monthly</p></div></div></div><div class="flex flex-col gap-[6px] items-start"><div class="flex items-center flex-col gap-[4px]"><img class="w-[24px] h-[24px] object-contain" src="/public/img/block-5.svg"><h4 class="text-[#2F2F2F] font-medium text-xs mb-0">Lorem ipsum</h4></div><p class="text-[#555555] text-[16px] font-normal m-0 pb-[8px] border-b border-[#ECECEC] w-full">Lorem ipsum dolor sit amet, consectetur, sed do eiusmod tempor.</p><div class="grid grid-cols-2 gap-[16px] w-full"><div class="flex flex-col items-center"><h3 class="text-center text-[#2F2F2F] text-[16px] mb-0">75%</h3><p class="text-center text-[#555555] text-[18px] mb-0">Reduced monthly</p></div><div class="flex flex-col items-center"><h3 class="text-center text-[#2F2F2F] text-[16px] mb-0">75%</h3><p class="text-center text-[#555555] text-[18px] mb-0">Reduced monthly</p></div></div></div></div>','/blocks/IMG-1726752455.jpeg',1,0,uid,'time',tid),('layout12','<div class="card-8 px-[8px] py-[12px] bg-white rounded flex flex-col gap-[8px]"><div class="rounded-[4px] border border-[#EDEDED] flex gap-[8px] items-center px-[12px] py-[10px]"><div class="w-[42px] h-[42px] rounded-full"><img src="/public/img/block-7.png" alt="" class="w-full h-full object-cover"></div><div class="flex flex-col items-start gap-[4px]"><h3 class="text-[#262626] text-xs mb-0 font-medium">Travis Orn</h3><p class="text-[#717171] font-light text-[16px] m-0">Software Engineer</p></div></div><div class="rounded-[4px] border border-[#EDEDED] flex gap-[8px] items-center px-[12px] py-[10px]"><div class="w-[42px] h-[42px] rounded-full"><img src="/public/img/block-7.png" alt="" class="w-full h-full object-cover"></div><div class="flex flex-col items-start gap-[4px]"><h3 class="text-[#262626] text-xs mb-0 font-medium">Travis Orn</h3><p class="text-[#717171] font-light text-[16px] m-0">Software Engineer</p></div></div></div>','/blocks/IMG-1726752540.jpeg',1,0,uid,'time',tid)

INSERT INTO tbl_blocks(title, block_content, cover_image, is_active , prime , created_by, created_on,tenant_id) VALUES('layout13','<div class="card-9 bg-white rounded grid grid-cols-2 gap-[8px] px-[8px] py-[12px]"><div class="relative pt-[36px]"><div class="px-[15px] pt-[38px] pb-[4px] bg-[#F9F8FB] rounded-[4px] flex flex-col gap-[2px] mb-[10px]"><h3 class="text-[#262626] text-xs font-medium mb-0 text-center">Eileen Walsh</h3><p class="text-[#717171] text-[16px] text-center">Software Engineer</p><div class="flex justify-center items-center gap-[16px]"><a href="#"><img src="/public/img/block-social1.svg" alt=""></a><a href="#"><img src="/public/img/block-social2.svg" alt=""></a><a href="#"><img src="/public/img/block-social3.svg" alt=""></a></div></div><div class="flex justify-center w-full absolute top-0 left-0"><img src="/public/img/block-8.png" class="w-[70px] h-[70px] rounded-full"></div></div><div class="relative pt-[36px]"><div class="px-[15px] pt-[38px] pb-[4px] bg-[#F9F8FB] rounded-[4px] flex flex-col gap-[2px] mb-[10px]"><h3 class="text-[#262626] text-xs font-medium mb-0 text-center">Eileen Walsh</h3><p class="text-[#717171] text-[16px] text-center">Software Engineer</p><div class="flex justify-center items-center gap-[16px]"><a href="#"><img src="/public/img/block-social1.svg" alt=""></a><a href="#"><img src="/public/img/block-social2.svg" alt=""></a><a href="#"><img src="/public/img/block-social3.svg" alt=""></a></div></div><div class="flex justify-center w-full absolute top-0 left-0"><img src="/public/img/block-8.png" class="w-[70px] h-[70px] rounded-full"></div></div></div>','/blocks/IMG-1726752710.jpeg',1,0,uid,'time',tid),('layout14','<div class="card-11 p-[8px] bg-white rounded grid grid-cols-3 gap-[8px]"><div class="flex flex-col gap-[6px] items-start"><div class="flex items-start flex-col gap-[4px]"><img class="w-[24px] h-[24px] object-contain" src="/public/img/block-11.svg"><h4 class="text-[#2F2F2F] font-medium text-xs mb-0">Lorem ipsum</h4></div><p class="text-[#555555] text-[16px] font-normal m-0 pb-[8px]  w-full">Loremipsum dolor sit amet, consectetur, sed do eiusmod tempor.</p></div><div class="flex flex-col gap-[6px] items-start"><div class="flex items-start flex-col gap-[4px]"><img class="w-[24px] h-[24px] object-contain" src="/public/img/block-11.svg"><h4 class="text-[#2F2F2F] font-medium text-xs mb-0">Lorem ipsum</h4></div><p class="text-[#555555] text-[16px] font-normal m-0 pb-[8px]  w-full">Lorem ipsum dolor sit amet, consectetur, sed do eiusmod tempor.</p></div><div class="flex flex-col gap-[6px] items-start"><div class="flex flex-col items-start gap-[4px]"><img class="w-[24px] h-[24px] object-contain" src="/public/img/block-11.svg"><h4 class="text-[#2F2F2F] font-medium text-xs mb-0">Lorem ipsum</h4></div><p class="text-[#555555] text-[16px] font-normal m-0 pb-[8px] w-full">Lorem ipsum dolor sit amet, consectetur, sed do eiusmod tempor.</p></div></div>','/blocks/IMG-1726752858.jpeg',1,0,uid,'time',tid),('layout15','<div class="card-13 p-[12px] bg-white rounded flex-col flex gap-[8px] items-start"><img class="w-[24px] h-[24px] object-contain" src="/public/img/block-11.svg"><h4 class="text-[#2F2F2F] font-medium text-xs mb-0 text-start">Lorem ipsum dolor Lorem ipsum dolor</h4><p class="text-[#555555] text-[16px] font-normal m-0 pb-[8px]  w-full text-start">Lorem ipsum dolor sit amet, consectetur, sed do eiusmod tempor.</p><a href="#" class="flex items-center justify-start text-[#10A37F] text-xs font-normal no-underline gap-[4px]">Read More<img src="/public/img/green-arrow.svg" alt=""></a></div>','/blocks/IMG-1726752982.jpeg',1,0,uid,'time',tid)

INSERT INTO tbl_blocks(title, block_content, cover_image, is_active , prime , created_by, created_on,tenant_id) VALUES('layout16','<div class="card-14 px-[8px] py-[16px] bg-white rounded flex gap-[8px] items-start"><img class="w-[32px] h-[32px] object-contain" src="/public/img/block-12.svg"><div class="flex flex-col gap-[8px] items-start"><h3 class="text-[#2F2F2F] text-[18px] font-medium m-0 text-start">Seamless integration with your schedule</h3><p class="text-[#555555] text-[16px] font-normal m-0 text-start">Guaranteed four hour overlap to ensure easier collaboration and communication. We adjust so you dont have to.</p></div></div>','/blocks/IMG-1726753113.jpeg',1,0,uid,'time',tid),('layout17','<div class="card-15 flex flex-col items-start rounded-[6px] bg-white "><img src="/public/img/block-13.png" class="w-full object-cover"><div class="flex gap-[8px]"><div class="p-[8px] "><div class="flex gap-[8px] items-center"><div class="flex flex-col gap-[4px] items-center"><div class="w-[38px] h-[38px] rounded-full"><img class="w-[38px] h-[38px] object-contain" src="/public/img/block-14.png"></div><h4 class="text-[#262626] text-[14px] m-0 font-normal whitespace-nowrap">Theodore Sipes</h4></div><div class="flex flex-col gap-[6px] items-start ps-[5px] border-l border-[#E5E7EB]"><p class="text-[#717171] text-[16px]">“Lorem ipsum dolor sit amet, consectetur, sed do eiusmod tempor. Lorem ipsum dolor sit amet.”</p></div></div></div></div></div>','/blocks/IMG-1726753198.jpeg',1,0,uid,'time',tid),('layout18','<div class="card-6 p-[12px] bg-white rounded flex gap-[8px]"><div class="flex items-start gap-[6px]"><div class="flex flex-col gap-[4px]"><div class="w-[24px] h-[24px]"><img src="/public/img/block-6.png" alt="" class="h-full w-full object-cover"></div><div class="w-[24px] h-[24px] object-contain"><img src="/public/img/block-6.png" alt="" class="h-full w-full object-cover"></div></div><div class="w-[80px] h-full"><img src="/public/img/block-6.png" alt="" class="w-full rounded-[2px] h-full"></div></div><div class="flex flex-col gap-[8px] items-start w-full"><div class="flex flex-col items-start gap-[4px]"><h5 class="text-xs font-normal mb-0 text-[#262626]">Lorem ipsum dolor sit</h5><div class="flex gap-[1px] items-center"><img src="img/reviewstar-full.svg" alt=""><img src="/public/img/reviewstar-full.svg" alt=""><img src="/public/img/reviewstar-full.svg" alt=""><img src="/public/img/reviewstar-full.svg" alt=""><img src="/public/img/reviewstar.svg" alt=""></div></div><p class="m-0 text-[#262626] text-xs font-normal">$ 15.00</p><div class="flex flex-col items-start gap-[4px]"><h5 class="text-[14px] font-normal mb-0 text-[#717171]">Color</h5><div class="flex items-center gap-[1px]"><div class="w-[20px] h-[20px] rounded-full p-[2px] border border-[#000000]"><div class="bg-[#000000] h-full rounded-full"></div></div><div class="w-[20px] h-[20px] rounded-full p-[2px]"><div class="bg-[#004DFF] h-full rounded-full"></div></div><div class="w-[20px] h-[20px] rounded-full p-[2px]"><div class="bg-[#3B8620] h-full rounded-full"></div></div></div></div><div class="flex flex-col items-start gap-[4px]"><h5 class="text-[14px] font-normal mb-0 text-[#717171]">Size</h5><div class="flex items-center gap-[8px]"><div class="h-[16px] flex items-center justify-center px-[6px] border border-[#EDEDED] text-[14px] text-[#262626] font-normal rounded-[4px] uppercase">s</div><div class="h-[16px] flex items-center justify-center px-[6px] border bg-[#EBEBEB] border-[#EDEDED] text-[14px] text-[#262626] font-normal rounded-[4px] uppercase">m</div><div class="h-[16px] flex items-center justify-center px-[6px] border border-[#EDEDED] text-[14px] text-[#262626] font-normal rounded-[4px] uppercase">l</div><div class="h-[16px] flex items-center justify-center px-[6px] border border-[#EDEDED] text-[14px] text-[#262626] font-normal rounded-[4px] uppercase">xl</div></div></div></div></div>','/blocks/IMG-1726753308.jpeg',1,0,uid,'time',tid)

INSERT INTO tbl_block_mstr_tags(tag_title, created_by ,created_on,tenant_id ) VALUES ('Default',uid,'time',tid)

-- block tags are inserted in the database based on the number of blocks inserted before in the tbl_blocks

INSERT INTO tbl_block_tags (block_id, tag_id, tag_name, created_by , created_on ,tenant_id ) VALUES (blid,tagid,'Default',uid,'time',tid)

-- block collections are inserted in the database based on the number of blocks inserted before in the tbl_blocks

INSERT INTO tbl_block_collections (user_id ,block_id,tenant_id) VALUES (uid,blid,tid)

